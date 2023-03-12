package util

import (
	"sync"

	"github.com/ikuraoo/fastdouyin/entity"
)

type Comment entity.Comment

func (Comment) TableName() string {
	return "comments"
}

type CommentDao struct {
}

var commentDao *CommentDao //DAO(DataAccessObject)模式
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

func (c *CommentDao) QueryByVId(vid int64) (*[]Comment, error) {
	var comments []Comment
	err := db.Where("vid = ?", vid).Find(&comments).Error
	if err != nil {
		Logger.Error("find comment by vid err:" + err.Error())
		return nil, err
	}
	return &comments, nil
}

func (c *CommentDao) CreateComment(content *Comment) (*Comment, error) {

	if err := db.Create(&content).Error; err != nil {
		Logger.Error("insert favourite err:" + err.Error())
		return nil, err
	}
	return content, nil
}
