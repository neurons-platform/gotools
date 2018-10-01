package utils

import (
	"reflect"
	"regexp"
	"testing"
)

func TestFormatToReg2(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		args args
		want *regexp.Regexp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatToReg2(tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatToReg2() = %v, want %v", got, tt.want)
			}
		})
	}
}
