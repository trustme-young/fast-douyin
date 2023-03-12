package entity

import "time"

type Favourite struct {
	Id          int64 `gorm:"column:id"`
	UId         int64 `gorm:"column:uid"`
	VId         int64 `gorm:"column:vid"`
	IsFavourite bool
	CreateTime  time.Time `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;default:null" json:"update_time"`
}
