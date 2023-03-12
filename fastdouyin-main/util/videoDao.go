package util

import (
	"sync"
	"time"

	"github.com/ikuraoo/fastdouyin/entity"

	"gorm.io/gorm"
)

type VideoDao struct{}

var videoDao *VideoDao //DAO(DataAccessObject)模式
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

type Videos entity.Video

func (*VideoDao) IncFavouriteCount(vid int64) error {
	// err := db.Model(Videos{}).Where("id = ?", vid).UpdateColumn("favourite_count", gorm.Expr("favourite_count + ?", 1)).Error
	err := db.Model(Videos{}).Where("id = ?", vid).Updates(map[string]interface{}{"favourite_count": gorm.Expr("favourite_count + ?", 1), "update_time": time.Now()}).Error
	if err != nil {
		Logger.Error("inc video favourite count error")
		return err
	}
	return nil
}

func (*VideoDao) DecFavouriteCount(vid int64) error {
	// err := db.Model(Videos{}).Where("id = ?", vid).UpdateColumn("favourite_count", gorm.Expr("favourite_count - ?", 1)).Error
	err := db.Model(Videos{}).Where("id = ?", vid).Updates(map[string]interface{}{"favourite_count": gorm.Expr("favourite_count - ?", 1), "update_time": time.Now()}).Error

	if err != nil {
		Logger.Error("dec video favourite count error")
		return err
	}
	return nil
}

func (*VideoDao) QueryVideoById(vid int64) (*Videos, error) {
	var videoList Videos
	err := db.Where("id = ?", vid).Find(&videoList).Error
	if err != nil {
		Logger.Error("find video by vid err:" + err.Error())
		return nil, err
	}
	return &videoList, nil
}

func (*VideoDao) IncCommentCount(vid int64) error {
	err := db.Model(Videos{}).Where("id = ?", vid).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		Logger.Error("inc video comment count error")
		return err
	}
	return nil
}

func (*VideoDao) DecCommentCount(vid int64) error {
	err := db.Model(Videos{}).Where("id = ?", vid).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	if err != nil {
		Logger.Error("dec video comment count error")
		return err
	}
	return nil
}
