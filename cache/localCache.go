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

// 检查key是否在缓存中，如果不存在就是设置key并设置过期时间
func CheckKey(k string, d time.Duration) bool {
	v := GetString(k)
	if len(v) > 0 {
		return true
	} else {
		SetString(k, "0", d)
	}
	return false
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
