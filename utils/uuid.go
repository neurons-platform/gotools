package utils

import (
	"github.com/satori/go.uuid"
)


//GetUUID 获取uuid
func GetUUID() string {
	v, _ := uuid.NewV4()
	id := uuid.Must(v, nil)
	return id.String()
}
