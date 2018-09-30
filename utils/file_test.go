package utils

import (
	"testing"
)

func TestReadAllFile(t *testing.T) {
     got := ReadAllFile("../linuxOS/centos.yaml")
     LogPrintln(got)
}
