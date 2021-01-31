package model

type Car struct {
	AM_MODEL
	CarRfidId   uint   `json:"car_rfid_id" gorm:"column:car_rfid_id"`
	MasterId    uint   `json:"master_id" gorm:"column:master_id"`
	Type        string `json:"type" gorm:"column:type"`
	PlateNumber string `json:"plate_number" gorm:"column:plate_number"`
	Color       string `json:"color" gorm:"column:color"`
	StartTime   string `json:"start_time" gorm:"column:start_time"`
	Flag        string `json:"flag" gorm:"column:flag"`
}
