package utils

import "testing"

func TestYamlStrToStruct(t *testing.T) {
	var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

	type T struct {
		A string
		B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
		}
	}
	tb := T{}
	YamlStrToStruct(data,&tb)
	LogPrintln(tb)
}
