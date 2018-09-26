package utils

import (
	"fmt"
	"os"
)

// 这里是对程序抛出异常的处理,直接打印异常
func Throw(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func Exit(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(2)
	}
}
