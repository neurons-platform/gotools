package file

import (
	"bufio"
	"fmt"
	"github.com/hpcloud/tail"
	"io"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

func Throw(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func DeleteFile(filePath string) bool {
	var err = os.Remove(filePath)
	if err != nil {
		return false
	}
	return true
}

func ReadAllFile(filePath string) string {
	b, e := ioutil.ReadFile(filePath)
	Throw(e)
	return string(b)
}

func WatchFile(fileName string, fun func(string)) {
	for {
		var seek *tail.SeekInfo
		seek = &tail.SeekInfo{}
		seek.Offset = 0
		seek.Whence = 2
		t, err := tail.TailFile(fileName, tail.Config{Follow: true, Location: seek})
		if Throw(err) {
			for line := range t.Lines {
				if line.Err == nil {
					fun(line.Text)
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func ReadFileToStringList(fileName string) []string {
	var stringList []string
	fi, _ := os.Open(fileName)
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		stringList = append(stringList, string(a))
	}
	return stringList

}

func statTimes(name string) (atime, mtime, ctime time.Time, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return
}

func GetFileModifyTime(path string) int64 {
	_, mtime, _, err := statTimes(path)
	if !Throw(err) {
		return 0
	}
	return GetMilliTimeStamp(mtime)
}

func GetFileCreateTime(path string) int64 {
	_, _, ctime, err := statTimes(path)
	if !Throw(err) {
		return 0
	}
	return GetMilliTimeStamp(ctime)
}

func CreateFile(path string) bool {

	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if !Throw(err) {
			return false
		}
		defer file.Close()
	}
	return true
}

func IsFileModifyTimeOlderThanNSecond(path string, n int) bool {
	if !FileExist(path) {
		return true
	}
	modifyTime := GetFileModifyTime(path)
	nowTimw := MilliTimeStamp()
	diffTime := (nowTimw - modifyTime) / 1000
	return diffTime > int64(n)
}

func WriteStringToFile(path string, content string) bool {
	if !FileExist(path) {
		CreateFile(path)
	}

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if !Throw(err) {
		return false
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if !Throw(err) {
		return false
	}

	err = file.Sync()
	if !Throw(err) {
		return false
	}
	return true
}
