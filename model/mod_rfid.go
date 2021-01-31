package model

type RFID struct {
	AM_MODEL
	RfidId    uint   `json:"rfid_id" gorm:"column:rfid_id"`
	CarRfidIf uint   `json:"car_rfid_id" gorm:"column:car_rfid_id"`
	Time      string `json:"time" gorm:"column:time"`
}
