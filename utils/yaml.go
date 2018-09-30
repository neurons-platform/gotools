package utils

import (
        "gopkg.in/yaml.v2"
)

func YamlStrToStruct(str string,t interface{}) bool {
	err := yaml.Unmarshal([]byte(str), t)
	return Throw(err)
}
