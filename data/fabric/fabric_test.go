package fabric

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/geraldkohn/elian/data/model"
)

func TestMarshalAndUnmarshalRecord(t *testing.T) {
	record := &model.Record{
		Uid: "812891238-1234",
		PatientUid: "123472134-14312",
		WriteStaffUid: "328294-2345-2345",
		ReadOnlyStaffUid: "132412-12341-1234",
		PhotoHash: "url-photo",
		DocumentHash: "url-document",
		Description: "test",
		LastChangeTime: "2022-1-1 09:09:09",
	}
	v := marshalRecord(record)
	log.Println(v)
	log.Println("-------------------------------------")
	newRecord := unmarshalRecord(v)
	log.Println(newRecord)
	if reflect.DeepEqual(*record, *newRecord) {
		log.Println("pass")
	} else {
		t.Errorf("wrong")
	}
}

func TestCreateAndFindRecord(t *testing.T) {
	record := &model.Record{
		Uid: "recordUid",
		PatientUid: "123472134-14312",
		WriteStaffUid: "328294-2345-2345",
		ReadOnlyStaffUid: "132412-12341-1234",
		PhotoHash: "url-photo",
		DocumentHash: "url-document",
		Description: "test",
		LastChangeTime: "2022-1-1 09:09:09",
	}
	recordRepo := NewRecordRepo()
	err := recordRepo.Create(context.Background(), record)
	if err != nil {
		log.Println(err)
	}

	newRecord, err := recordRepo.FindByUid(context.Background(), record.Uid)
	if err != nil {
		log.Panicln(err)
	}
	if reflect.DeepEqual(*record, *newRecord) {
		log.Println("pass")
	} else {
		t.Errorf("wrong")
	}

	str, err := recordRepo.FindWriteStaffUidByUid(context.Background(), record.Uid)
	if err != nil {
		log.Panicln(err)
	}
	if reflect.DeepEqual(str, record.WriteStaffUid) {
		log.Println("pass")
	} else {
		t.Errorf("wrong")
	}
}

func TestUpdateRecord(t *testing.T) {
	record := &model.Record{
		Uid: "recordUid",
		PatientUid: "123472134-14312",
		WriteStaffUid: "328294-2345-2345",
		ReadOnlyStaffUid: "132412-12341-1234",
		PhotoHash: "url-photo",
		DocumentHash: "url-document",
		Description: "test",
		LastChangeTime: "2022-1-1 09:09:09",
	}
	recordRepo := NewRecordRepo()
	err := recordRepo.Create(context.Background(), record)
	if err != nil {
		log.Println(err)
	}

	err = recordRepo.UpdateDescriptionByUid(context.Background(), record.Uid, "newtest")
	if err != nil {
		log.Println(err)
	}

	newRecord, err := recordRepo.FindByUid(context.Background(), record.Uid)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(newRecord)
	if newRecord.Description == "newtest" {
		log.Println("pass")
	} else {
		t.Error()
	}
}