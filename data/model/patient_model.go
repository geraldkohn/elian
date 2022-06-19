package model

type Patient struct {
	Uid           string `json:"uid" gorm:"primary_key"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	IdCardNumber  string `json:"id_card_number"`
	RegisterTime  string `json:"register_time"`
	LastLoginTime string `json:"last_login_time"`
	Token         string `json:"token"`
}

func (Patient) TableName() string {
	return "patient"
}
