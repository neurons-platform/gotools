package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

// 序列化 golang的对象(struct map 等)
func Obj2str(m interface{}) string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	Throw(err)
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

// 反序列化
func Str2obj(str string, m interface{}) {
	by, err := base64.StdEncoding.DecodeString(str)
	Throw(err)
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(m)
	Throw(err)
}
