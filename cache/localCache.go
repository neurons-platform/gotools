package cache

import (
	CH "github.com/patrickmn/go-cache"
	"time"
)

var C = CH.New(5*time.Minute, 10*time.Minute)


func Set(k string, x interface{},d time.Duration) {
	C.Set(k, x, d)
}


func Delete(k string) {
	C.Delete(k)
}

func SetString(k string, x string , d time.Duration) {
	C.Set(k, x, d)
}

func GetString(k string) string {
	r := ""
	x, ok := C.Get(k)
	if ok {
		r = x.(string)
	}
	return r
}
