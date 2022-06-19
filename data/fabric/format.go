package fabric

import (
	"encoding/json"
	"log"

	"github.com/geraldkohn/elian/data/model"
)

func marshalRecord(record *model.Record) (v string) {
	b, err := json.Marshal(record)
	if err != nil {
		log.Println("record 格式化错误")
	}
	v = string(b)
	return
}

func unmarshalRecord(v string) *model.Record {
	b := []byte(v)
	r := &model.Record{}
	err := json.Unmarshal(b, &r)
	if err != nil {
		log.Println("record 解析错误")
	}
	return r
}
