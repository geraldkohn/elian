package model

/**
由于数据库不支持数组类型的变量, 所以WriteStaffUid和ReadOnlyStaffUid只能存String类型
暂定这这两个字段的格式是: uid1&uid2&uid3.....&uidn
*/

type Record struct {
	Uid string `json:"uid" gorm:"primary_key"`

	PatientUid       string `json:"patient_uid"`         //这条记录对应的患者的uid
	WriteStaffUid    string `json:"staff_uid"`           //医生的uid --- 对这个记录可以读和写
	ReadOnlyStaffUid string `json:"read_only_staff_uid"` //医生的uid --- 对这个记录只能写

	PhotoHash    string `json:"photo_hash"`    //图片资源的唯一标识
	DocumentHash string `json:"document_hash"` //文档资源的唯一标识

	LastChangeTime string `json:"last_change_time"` //此记录的创建时间
	Description  string `json:"description"`   //病情的简要描述
}

func (Record) TableName() string {
	return "record"
}
