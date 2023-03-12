package service

import (
	"errors"
	"time"

	"github.com/ikuraoo/fastdouyin/util"
)

type FavouriteActionFlow struct {
	UId int64
	VId int64
}

func FavouriteAction(userId int64, videoId int64) error {
	return NewFavouriteActionFlow(userId, videoId).Do()
}

func NewFavouriteActionFlow(userId int64, videoId int64) *FavouriteActionFlow {
	return &FavouriteActionFlow{
		UId: userId,
		VId: videoId,
	}
}

func (f *FavouriteActionFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	err := f.action()
	if err != nil {
		return err
	}
	return nil
}

func (f *FavouriteActionFlow) checkParam() error {
	return nil
}

func (f *FavouriteActionFlow) action() error {
	isFavourite, err := util.NewFavouriteDaoInstance().QueryByVIdAndUId(f.VId, f.UId)

	if err != nil {
		//没有找到
		if err := util.NewFavouriteDaoInstance().CreateFavourite(&util.Favourites{
			UId:         f.UId,
			VId:         f.VId,
			IsFavourite: true,
			CreateTime:  time.Now(),
			UpdateTime:  time.Now(),
		}); err != nil {
			return err
		}

		err := util.NewVideoDaoInstance().IncFavouriteCount(f.VId)
		if err != nil {
			return err
		}

		return nil
	}
	if isFavourite == true {
		err := util.NewVideoDaoInstance().DecFavouriteCount(f.VId)
		if err != nil {
			return err
		}
	} else {
		err := util.NewVideoDaoInstance().IncFavouriteCount(f.VId)
		if err != nil {
			return err
		}
	}
	err = util.NewFavouriteDaoInstance().UpdateIsFavourite(f.VId, f.UId, !isFavourite)

	if err != nil {
		return errors.New("修改失败")
	}
	return nil
}
