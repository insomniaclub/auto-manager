package model

type Manager struct {
	AM_MODEL
	Nickname string `json:"nickname" gorm:"column:nickname"`
	Password string `json:"-" gorm:"column:password"`
}
