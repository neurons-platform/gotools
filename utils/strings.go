package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Str2Int(str string) int {
	i, err := strconv.Atoi(str)
	Throw(err)
	return i

}

func Str2Int64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	Throw(err)
	return i

}

func Str2Float64(str string) float64 {
	i, err := strconv.ParseFloat(str, 64)
	Throw(err)
	return i

}

func Float64toStr(i float64) string {
	return strconv.FormatFloat(i, 'f', 6, 64)
}

func Int2Str(i int) string {
	s := strconv.Itoa(i)
	return s
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

func Int32toStr(i int32) string {
	s := strconv.FormatInt(int64(i), 10)
	return s
}

func Int64toStr(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}

func BoolToStr(b bool) string {
	return strconv.FormatBool(b)
}

func StrToBool(s string) (bool, error) {
	switch s {
	case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "y", "ON", "on", "On":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "n", "OFF", "off", "Off":
		return false, nil
	}
	return false, fmt.Errorf("parsing \"%s\": invalid syntax", s)
}

func StartWith(str string, pre string) bool {
	return strings.HasPrefix(str, pre)
}
