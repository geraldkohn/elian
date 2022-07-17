package record

import (
	"log"
	"testing"
)

func TestMarshal(t *testing.T) {
	str, err := marshal("", "adfasdfasdfdsaf")
	if err != nil {
		log.Println(err)
	}
	uids, err := unmarshal(str)
	if err != nil {
		log.Println(err)
	}
	for _, s := range uids {
		log.Println(s)
	}

	nextStr1, err := marshal(str, "test2")
	if err != nil {
		log.Println(err)
	}
	nextStr2, err := marshal(nextStr1, "test3")
	if err != nil {
		log.Println(err)
	}
	nextUids, err := unmarshal(nextStr2)
	if err != nil {
		log.Println(err)
	}
	log.Println("----------------")
	log.Println(nextStr2)
	log.Println("begin:------------------")
	for _, s := range nextUids {
		log.Println(s)
	}
}
