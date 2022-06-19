package model

type Agency struct {
	Uid           string `json:"uid" gorm:"primary_key"`
	License       string `json:"license"`
	RegisterTime  string `json:"register_time"`
	LastLoginTime string `json:"last_login_time"`
	Token         string `json:"token"`
}

func (Agency) TableName() string {
	return "agency"
}
