package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"text/template"
)

func ParserMapToString(m map[string]string, str string) string {
	tmp := StrToTemplate(str)
	t := template.Must(template.New("tmp").Parse(tmp))
	var buf bytes.Buffer
	t.Execute(&buf, m)
	return buf.String()
}

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return Int2Str(v)
	case int32:
		return Int32toStr(v)
	case int64:
		return Int64toStr(v)
	case float64:
		return Int64toStr(int64(v))
		// Add whatever other types you need
	default:
		fmt.Printf("%s", v)
		return "no"
	}
}

func ParserJsonStrToTemplateStr(js string, tpl string) (string, bool) {
	m := map[string]interface{}{}
	if err := json.Unmarshal([]byte(js), &m); err != nil {
		log.Fatal(err)
		return "", false
	}

	t := template.New("tmp")
	t.Funcs(template.FuncMap{"tostring": ToString})
	t.Parse(tpl)
	// t := template.Must(template.New("tmp").Parse(tpl))

	var buf bytes.Buffer
	t.Execute(&buf, m)
	return buf.String(), true
}
