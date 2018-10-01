package utils

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func GetLocalIp() string {
	var ip string
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}

		}
	}
	return ip
}

func GetOsType() string {
	var osType = ""
	if runtime.GOOS == "linux" {
		osType = "linux"
	}

	if runtime.GOOS == "windows" {
		osType = "windows"
	}
	return osType
}

