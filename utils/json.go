package utils

import (
	"bytes"
	"encoding/json"
	"log"
	SJ "github.com/bitly/go-simplejson"
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

func JsonStrToJsonSJ(str string) ( *SJ.Json,error) {
	js, err := SJ.NewJson([]byte(str))
	return js,err
}
