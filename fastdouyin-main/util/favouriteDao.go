package util

import (
	"sync"
	"time"

	"github.com/ikuraoo/fastdouyin/entity"
)

type Favourites entity.Favourite

func (Favourites) TableName() string {
	return "favourites"
}

type FavouriteDao struct {
}

var favouriteDao *FavouriteDao //DAO(DataAccessObject)模式
var favouriteOnce sync.Once

func NewFavouriteDaoInstance() *FavouriteDao {
	favouriteOnce.Do(
		func() {
			favouriteDao = &FavouriteDao{}
		})
	return favouriteDao
}

// QueryByVIdAndUId 如果存在，则返回fav.IsFavourite，否则返回err
func (*FavouriteDao) QueryByVIdAndUId(vid int64, uid int64) (bool, error) {
	var fav *Favourites
	// fav = new(Favourites)
	err := db.Where("uid = ?", uid).Where("vid = ?", vid).First(&fav).Error //链式操作

	if err != nil {
		Logger.Error("find favourite by vid and uid err:" + err.Error())
		return false, err
	}

	return fav.IsFavourite, nil
}

// QueryByUId 如果存在，返回列表，否则返回空，报错则返回err
func (*FavouriteDao) QueryByUId(uid int64) (*[]Favourites, error) {
	var fav []Favourites
	err := db.Where("uid = ?", uid).Find(&fav).Error
	if err != nil {
		Logger.Error("find favourite by id err:" + err.Error())
		return nil, err
	}
	return &fav, nil
}

// UpdateIsFavourite 若点赞了，就取消；若没有，则点赞
func (f *FavouriteDao) UpdateIsFavourite(vid int64, uid int64, IsFavourite bool) error {
	// err := db.Model(Favourites{}).Where("uid = ?", uid).Where("vid = ?", vid).Update("is_favourite", IsFavourite).Error
	// db.Model(Favourites{}).Where("uid = ?", uid).Where("vid = ?", vid).Update("update_time", time.Now())
	err := db.Model(Favourites{}).Where("uid = ?", uid).Where("vid = ?", vid).Updates(map[string]interface{}{"is_favourite": IsFavourite, "update_time": time.Now()}).Error

	if err != nil {
		return err
	}

	return nil
}

func (f *FavouriteDao) CreateFavourite(fav *Favourites) error {
	if err := db.Create(&fav).Error; err != nil {
		println("insert favourite err:")
		Logger.Error("insert favourite err:" + err.Error())
		return err
	}
	return nil
}

// QueryByVIdAndUId 如果存在，则返回fav.IsFavourite，否则返回err
func (*FavouriteDao) QueryByUIdAndIsFavourite(uid int64) (*[]int64, error) {
	var favs []Favourites
	var vidList []int64
	err := db.Where("uid = ?", uid).Where("is_favourite = ?", true).Find(&favs).Error //链式操作

	if err != nil {
		Logger.Error("find vid by uid and isfavourite err:" + err.Error())
		return nil, err
	}
	for _, fav := range favs {
		vidList = append(vidList, fav.VId)
	}
	return &vidList, nil
}
