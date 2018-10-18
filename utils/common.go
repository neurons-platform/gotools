package utils

import (
	"bytes"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"reflect"
)

func SumIntArray(a []int) int {
	var r = 0
	for _, v := range a {
		r = r + v
	}
	return r
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

// 返回在a中有 在b中没有的元素
func DiffArr(a []string, b []string) []string {
	var r []string
	for _, v := range a {
		have, _ := Contain(v, b)
		if !have {
			r = append(r, v)
		}
	}
	return r
}

func Call(func_name interface{}, params ...interface{}) func() {
	f := func() {
		var fc = reflect.ValueOf(func_name)
		in := make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
		fc.Call(in)
	}
	return f
}


func UniqueStrList(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func MaxInFloat64Aarry(max float64,arr []float64) float64 {
	for _, e := range arr {
		if max < e {
			max = e
		}
	}
	return max
}
