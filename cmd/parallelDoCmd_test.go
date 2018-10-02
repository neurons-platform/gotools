package cmd

import (
	// "reflect"
	"testing"
	"fmt"
)

func TestParallelDo(t *testing.T) {
	cmds := []CMD{
		{IP:"172.30.42.3",User:"root",Port:22,Command:"ps -ef |grep java",Password:"123456",CMDType:SSH,TimeOut:5},
		{IP:"172.30.42.2",User:"root",Port:22,Command:"date -R ; sleep 1",Password:"123456",CMDType:SSH,TimeOut:5},
		{IP:"172.30.42.3",User:"root",Port:22,Command:"date -R ; sleep 1",Password:"123456",CMDType:SSH,TimeOut:5},
		{IP:"172.30.42.4",User:"root",Port:22,Command:"date -R ; sleep 1",Password:"123456",CMDType:SSH,TimeOut:5},
		// {IP:"172.30.42.4",User:"root",Port:22,Command:"uptime",Password:"12345",CMDType:SSH,TimeOut:5},
	}
	got := ParallelDo(cmds,30)
	fmt.Println(got)
}
