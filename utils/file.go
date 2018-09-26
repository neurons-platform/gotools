package utils

import (
	"os"
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
