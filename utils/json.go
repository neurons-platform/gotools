package utils

import (
	"bytes"
	"encoding/json"
	SJ "github.com/bitly/go-simplejson"
	"log"
)

func StructToJsonString(st interface{}) string {
	j, _ := json.Marshal(st)
	return string(j)
}

func StructToJsonStringNotEscapHTML(st interface{}) string {

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(st); err != nil {
		log.Println(err)
	}
	return buf.String()
}

func JsonStringToStruct(str string, s interface{}) {
	json.Unmarshal([]byte(str), &s)
}

func JsonStringToMap(str string) map[string]interface{} {
	mp := make(map[string]interface{})
	if err := json.Unmarshal([]byte(str), &mp); err != nil {
		Throw(err)
		return nil
	}
	return mp
}

func JsonStrToJsonSJ(str string) (*SJ.Json, error) {
	js, err := SJ.NewJson([]byte(str))
	return js, err
}

func InterfaceToJsonStr(i interface{}) string {
	jsonStr, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	str := string(jsonStr[:])
	return str
}

type JS struct {
	Jstr string
}

func (this JS) String() string {
	return this.Jstr
}

func (this JS) Get(key string) JS {
	js := JS{
		Jstr: getJsonStrFromJsonStr(this.Jstr, key),
	}
	return js
}

func getJsonStrFromJsonStr(js string, key string) string {
	var raw map[string]interface{}
	json.Unmarshal([]byte(js), &raw)
	r := ""
	if val, ok := raw[key]; ok {
		out, _ := json.Marshal(val)
		r = string(out)
	}
	return r
}
