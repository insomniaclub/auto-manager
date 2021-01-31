package model

type Master struct {
	AM_MODEL
	IdNumber     string `json:"id_number" gorm:"column:id_number"`
	Name         string `json:"name" gorm:"column:name"`
	Nickname     string `json:"nickname" gorm:"column:nickname"`
	OpenId       string `json:"open_id" gorm:"column:open_id"`
	MobileNumber string `json:"mobile_number" gorm:"column:mobile_number"`
	RegisterTime string `json:"register_time" gorm:"column:register_time"`
}
