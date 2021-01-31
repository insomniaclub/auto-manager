package service

import (
	"auto-manager/global"
	"auto-manager/model"
)

func AddPosition(position *model.Position) (*model.Position, error) {
	err := global.AM_DB.Table("position_desc").Create(position).Error
	return position, err
}

func DeletePosition(id uint) (*model.Position, error) {
	var position model.Position
	err := global.AM_DB.Table("position_desc").Where("id = ?", id).Delete(&position).Error
	return &position, err
}

func ModifyPosition(p *model.Position) (*model.Position, error) {
	err := global.AM_DB.Table("position_desc").Where("id = ?", p.ID).Update("rfid_id", p.RfidId).Update("position", p.Position).Error
	return p, err
}

func GetPositionList() (*[]model.Position, error) {
	var positions []model.Position
	err := global.AM_DB.Table("position_desc").Find(&positions).Error
	return &positions, err
}
