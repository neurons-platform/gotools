package cmd

import (
	// "reflect"
	"testing"
	"fmt"
)

func TestDoCmd(t *testing.T) {
	type args struct {
		cmd CMD
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{
			args{
				cmd: CMD{IP:"172.30.42.2",User:"root",Port:22,Command:"date -R ; sleep 1",Password:"123456",CMDType:SSH},
			},
		},
		{
			args{
				cmd: CMD{IP:"172.30.42.4",User:"root",Port:22,Command:"date -R ; sleep 1",Password:"123456",CMDType:SSH},
			},
		},
		{
			args{
				cmd: CMD{IP:"172.30.42.3",User:"root",Port:22,Command:"date -R ; sleep 1",Password:"123456",CMDType:SSH},
			},
		},
	}
	for _, tt := range tests {
		got,ok := DoCmd(tt.args.cmd)
		fmt.Println(ok)
		fmt.Println(got.Out.GetStdOut())
	}
}
