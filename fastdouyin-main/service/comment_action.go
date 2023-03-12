package service

import (
	"github.com/ikuraoo/fastdouyin/util"

	"time"
)

type CommentActionFlow struct {
	UId     int64
	VId     int64
	Content string
}

func CommentAction(userId int64, videoId int64, content string) (*CommentProto, error) {
	return NewCommentActionFlow(userId, videoId, content).Do()
}

func NewCommentActionFlow(userId int64, videoId int64, content string) *CommentActionFlow {
	return &CommentActionFlow{
		UId:     userId,
		VId:     videoId,
		Content: content,
	}
}

func (c *CommentActionFlow) Do() (*CommentProto, error) {
	if err := c.checkParam(); err != nil {
		return nil, err
	}
	commentproto, err := c.action()
	if err != nil {
		return nil, err
	}
	return commentproto, nil
}

func (c *CommentActionFlow) checkParam() error {
	return nil
}

func (c *CommentActionFlow) action() (*CommentProto, error) {
	newcomment, err := util.NewCommentDaoInstance().CreateComment(&util.Comment{
		VId:        c.VId,
		UId:        c.UId,
		Content:    c.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsDeleted:  false,
	})
	//作者信息
	author, _ := util.NewUserDaoInstance().QueryUserById(c.UId)
	authorproto := &AuthorProto{
		Id:            author.Id,
		Name:          author.Name,
		FollowCount:   author.FollowCount,
		FollowerCount: author.FollowerCount,
		IsFollow:      true, //待完善
	}
	commentproto := CommentProto{
		Id:         newcomment.Id,
		User:       *authorproto,
		Content:    newcomment.Content,
		CreateDate: newcomment.CreateTime.GoString(),
	}

	if err != nil {
		return nil, err
	}
	//增加评论数
	err = util.NewVideoDaoInstance().IncCommentCount(c.VId)
	if err != nil {
		return nil, err
	}
	return &commentproto, nil
}
