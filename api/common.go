package main

//用户相关

type patientRegisterReq struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	IdCardNumber string `json:"id_card_number"`
}

type patientLoginReq struct {
	IdCardNumber string `json:"id_card_number"`
	Password     string `json:"password"`
}

type staffRegisterReq struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	Hospital     string `json:"hospital"`
	Department   string `json:"department"`
	JobNumber    string `json:"job_number"`
	IdCardNumber string `json:"id_card_number"`
}

type staffLoginReq struct {
	IdCardNumber string `json:"id_card_number"`
	Password     string `json:"password"`
}

type agencyRegisterReq struct {
	License string `json:"license"`
}

type agencyLoginReq struct {
	License string `json:"license"`
}

type registerAndLoginRes struct {
	Token     string `json:"token"`
	Status    int64  `json:"status"`
	Msg       string `json:"msg"`
	Attribute string `json:"attribute"`
}

//病历相关
//公共字段

type statusAndMsgRes struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
}

type recordRes struct {
	PatientIdCardNumber string `json:"patientIdCardNumber"`
	PatientName         string `json:"patientName"`
	PhotoHash           string `json:"photoHash"`
	DocumentHash        string `json:"documentHash"`
	LastChangeTime      string `json:"lastChangeTime"`
	Description         string `json:"description"`
	RecordUid           string `json:"recordUid"`
}

//私有字段

type staffCreateRecordReq struct {
	PatientIdCard string `json:"patientIdCardNumber"`
	PhotoHash     string `json:"photoHash"`
	DocumentHash  string `json:"documentHash"`
	Description   string `json:"description"`
}

type staffCreateRecordRes struct {
	StatusAndMsgRes statusAndMsgRes
}

type staffQueryRecordReq struct {
	PatientIdCardNumber string `json:"patientIdCardNumber"`
}

type staffQueryRecordRes struct {
	RecordResArray  []recordRes
	StatusAndMsgRes statusAndMsgRes
}

type patientQueryRecordReq struct {
	PatientIdCardNumber string `json:"patientIdCardNumber"`
}

type patientQueryRecordRes struct {
	RecordResArray  []recordRes
	StatusAndMsgRes statusAndMsgRes
}

type agencyQueryRecordReq struct {
	PatientIdCardNumber string `json:"patientIdCardNumber"`
}

type agencyQueryRecordRes struct {
	RecordResArray  []recordRes
	StatusAndMsgRes statusAndMsgRes
}

type staffSetReadAndWritePermissionReq struct {
	Hospital   string `json:"hospital"`
	Department string `json:"department"`
	Name       string `json:"name"`
	RecordUid  string `json:"recordUid"`
}

type staffSetReadAndWritePermissionRes struct {
	StatusAndMsgRes statusAndMsgRes
}

type staffSetReadOnlyPermissionReq struct {
	Hospital   string `json:"hospital"`
	Department string `json:"department"`
	Name       string `json:"name"`
	RecordUid  string `json:"recordUid"`
}

type agencySetReadOnlyPermissionRes struct {
	StatusAndMsgRes statusAndMsgRes
}

type agencySetReadAndWritePermissionReq struct {
	Hospital   string `json:"hospital"`
	Department string `json:"department"`
	Name       string `json:"name"`
	RecordUid  string `json:"recordUid"`
}

type agencySetReadAndWritePermissionRes struct {
	StatusAndMsgRes statusAndMsgRes
}

type agencySetReadOnlyPermissionReq struct {
	Hospital   string `json:"hospital"`
	Department string `json:"department"`
	Name       string `json:"name"`
	RecordUid  string `json:"recordUid"`
}

type staffSetReadOnlyPermissionRes struct {
	StatusAndMsgRes statusAndMsgRes
}

type staffUpdatePhotoReq struct {
	PhotoHash  string `json:"photoHash"`
	RecordUid string `json:"recordUid"`
}

type staffUpdatePhotoRes struct {
	StatusAndMsgRes statusAndMsgRes
}

type staffUpdateDocumentReq struct {
	DocumentHash string `json:"documentHash"`
	RecordUid   string `json:"recordUid"`
}

type staffUpdateDocumentRes struct {
	StatusAndMsgRes statusAndMsgRes
}

type staffUpdateDescriptionReq struct {
	Description string `json:"description"`
	RecordUid   string `json:"recordUid"`
}

type staffUpdateDescriptionRes struct {
	StatusAndMsgRes statusAndMsgRes
}
