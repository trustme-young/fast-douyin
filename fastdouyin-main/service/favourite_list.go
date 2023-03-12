package service

import (
	"github.com/ikuraoo/fastdouyin/entity"
	"github.com/ikuraoo/fastdouyin/util"
)

type Video entity.Video

func FavoriteList(uid int64) ([]*VideoProto, error) {
	_, err := util.NewUserDaoInstance().QueryUserById(uid)
	if err != nil {
		return nil, err
	}
	vids, err := util.NewFavouriteDaoInstance().QueryByUIdAndIsFavourite(uid)
	if err != nil {
		return nil, err
	}

	var videoprotoList []*VideoProto
	println(len(*vids))
	for _, vid := range *vids {
		videoList, err := util.NewVideoDaoInstance().QueryVideoById(vid)
		if err != nil {
			return nil, err
		}
		uid := videoList.UId
		//作者信息
		author, _ := util.NewUserDaoInstance().QueryUserById(uid)
		authorproto := &AuthorProto{
			Id:            author.Id,
			Name:          author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      true, //待完善
		}
		videoproto := VideoProto{
			Id:             videoList.Id,
			Author:         *authorproto,
			PlayUrl:        videoList.PlayUrl,
			CoverUrl:       videoList.CoverUrl,
			FavouriteCount: videoList.FavouriteCount,
			CommentCount:   videoList.CommentCount,
			IsFavourite:    true,
			Title:          videoList.Title,
		}
		videoprotoList = append(videoprotoList, &videoproto)
	}
	return videoprotoList, nil

}
