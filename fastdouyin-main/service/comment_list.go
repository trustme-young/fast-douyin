package service

import (
	"github.com/ikuraoo/fastdouyin/util"
)

func CommentList(vid int64) ([]*CommentProto, error) {

	commentList, err := util.NewCommentDaoInstance().QueryByVId(vid)
	if err != nil {
		return nil, err
	}
	var protoCommentList []*CommentProto
	for _, comment := range *commentList {
		user, err := util.NewUserDaoInstance().QueryUserById(comment.UId)
		if err != nil {
			return nil, err
		}
		demoUser := &AuthorProto{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
		month := comment.CreateTime.Format("01")
		date := comment.CreateTime.Format("02")

		protoCommentList = append(protoCommentList, &CommentProto{
			Id:         comment.Id,
			User:       *demoUser,
			Content:    comment.Content,
			CreateDate: month + "-" + date,
		})
	}
	return protoCommentList, nil
}
