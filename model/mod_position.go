package model

type Position struct {
	AM_MODEL
	RfidId   uint   `json:"rfid_id" gorm:"column:rfid_id"`
	Position string `json:"position" gorm:"column:position"`
}
