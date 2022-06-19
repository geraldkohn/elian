package uid

import (
	"github.com/google/uuid"
)

func NewUid() string {
	uid, err := uuid.NewRandom()
	if err != nil {
		panic("uid 生成失败")
	}
	return uid.String()
}
