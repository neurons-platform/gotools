package utils

import (
	"os"
	"io/ioutil"
)

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
