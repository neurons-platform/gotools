package cache

import (
	"testing"
	"time"
	"fmt"
)

func TestSetString(t *testing.T) {
	SetString("name", "neurons-platform", time.Minute*1)
	got := GetString("name")
	got = GetString("nam")
	fmt.Println(got)
}
