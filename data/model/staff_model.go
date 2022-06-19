package model

/**
hospital, department, name 确定一个医生
jobNumber, name 用来进行医生的认证
*/

type Staff struct {
	Uid           string `json:"uid" gorm:"primary_key"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	Hospital      string `json:"hospital"`
	Department    string `json:"department"`
	JobNumber     string `json:"job_number"`
	IdCardNumber  string `json:"id_card_number"`
	RegisterTime  string `json:"register_time"`
	LastLoginTime string `json:"last_login_time"`
	Token         string `json:"token"`
}

func (Staff) TableName() string {
	return "staff"
}
