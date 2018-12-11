package utils

import (
	"os"
)

func CreateDirIfNotExit(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if !Throw(err) {
			return false
		}
	}
	return true
}
