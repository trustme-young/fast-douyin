package entity

import (
	"time"
)

type Video struct {
	Id             int64     `gorm:"column:id" redis:"id"`
	UId            int64     `gorm:"column:uid" redis:"uid"`
	PlayUrl        string    `gorm:"column:play_url" redis:"play_url"`
	CoverUrl       string    `gorm:"column:cover_url" redis:"cover_url"`
	CommentCount   int64     `gorm:"column:comment_count" redis:"comment_count"`
	FavouriteCount int64     `gorm:"column:favourite_count" redis:"favorite_count"`
	Title          string    `gorm:"column:title" redis:"title"`
	CreateTime     time.Time `gorm:"column:create_time" redis:"-"`
	UpdateTime     time.Time `gorm:"column:update_time" redis:"-"`
	IsDeleted      bool      `gorm:"column:is_deleted" redis:"-"`
}
