package utils

import "testing"

func TestGetLocalIp(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLocalIp(); got != tt.want {
				t.Errorf("GetLocalIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOsType(t *testing.T) {
     got := GetOsType()
     LogPrintln(got)
}
