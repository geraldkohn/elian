package model

type PatientuidToRecorduid struct {
	Uid        string `json:"uid" gorm:"primary_key"`
	PatientUid string `json:"patient_uid"`
	RecordUid  string `json:"record_uid"`
}

func (PatientuidToRecorduid) TableName() string {
	return "patientuid_to_recorduid"
}