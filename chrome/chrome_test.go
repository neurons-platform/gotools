package chrome

import "testing"

func TestSavePageAsPng(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SavePageAsPng(tt.args.url); got != tt.want {
				t.Errorf("SavePageAsPng() = %v, want %v", got, tt.want)
			}
		})
	}
}
