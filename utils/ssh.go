package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Out struct {
	Stdout string
	Stderr string
}

func (o Out) GetStdOut() string {
	return o.Stdout
}

func (o Out) GetStdErr() string {
	return o.Stderr
}

func Scp(user, password, hostname string, port string, src, dst string) map[string]Out {
	var (
		auth         []ssh.AuthMethod
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	data := make(map[string]Out)
	if client, err = ssh.Dial("tcp", hostname+":"+port, clientConfig); err != nil {
		data[hostname] = Out{"error", "error"}
		return data
	}

	defer client.Close()

	if session, err = client.NewSession(); err != nil {
		data[hostname] = Out{"error", "error"}
		return data
	}
	defer session.Close()

	var stdout string
	var stderr string
	file, err := os.Open(src)
	if err != nil {
		//fmt.Println("打开文件失败:", err)
		stderr = "没找到文件"
		data[hostname] = Out{stdout, stderr}
		return data

	}
	info, _ := file.Stat()
	size := info.Size()
	filename := filepath.Base(src)
	dirname := strings.Replace(dst, "\\", "/", -1)

	go func() {
		w, _ := session.StdinPipe()
		fmt.Fprintln(w, "C0644", size, filename)
		io.CopyN(w, file, size)
		fmt.Fprint(w, "\x00")
		w.Close()
	}()
	if err := session.Run(fmt.Sprintf("/usr/bin/scp -qrt %s", dirname)); err != nil {
		stderr = "执行scp命令失败"
		session.Close()
		//return
	} else {
		stdout = "发送成功"
		session.Close()
	}

	data[hostname] = Out{stdout, stderr}

	session.Close()
	client.Close()
	return data
}

func ExecuteCmd(hostname, cmd, user, password, port string) map[string]Out {
	var (
		auth []ssh.AuthMethod
		//addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	data := make(map[string]Out)
	if client, err = ssh.Dial("tcp", hostname+":"+port, clientConfig); err != nil {
		error := err.Error()
		fmt.Println(error)
		data[hostname] = Out{"error", "error"}
		return data
	}
	if session, err = client.NewSession(); err != nil {
		fmt.Print("2\n")
		data[hostname] = Out{"error", "error"}
		return data
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	session.Run(cmd)

	data[hostname] = Out{stdoutBuf.String(), stderrBuf.String()}
	session.Close()
	client.Close()
	return data
}
