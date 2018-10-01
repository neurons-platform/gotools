package cache

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    *Cache
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithEvict(t *testing.T) {
	type args struct {
		size      int
		onEvicted func(key interface{}, value interface{})
	}
	tests := []struct {
		name    string
		args    args
		want    *Cache
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithEvict(tt.args.size, tt.args.onEvicted)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithEvict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithEvict() = %v, want %v", got, tt.want)
			}
		})
	}
}
