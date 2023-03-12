package util

import (
	"sync"

	"github.com/ikuraoo/fastdouyin/entity"
	"gorm.io/gorm"
)

type User entity.User

func (User) TableName() string {
	return "users"
}

type UserDao struct {
}

var userDao *UserDao //DAO(DataAccessObject)模式
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) CreateUser(user *User) error {

	if err := db.Create(user).Error; err != nil {
		Logger.Error("insert user err:" + err.Error())
		return err
	}
	return nil
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil

}

func (*UserDao) QueryUserByName(name string) (*User, error) {
	var user User
	err := db.Where("name = ?", name).Find(&user).Error
	if err != nil {
		Logger.Error("find user by name err:" + err.Error())
		return nil, err
	}
	return &user, nil

}

func (*UserDao) IncUserFollow(uid int64) error {
	err := db.Model(User{}).Where("id = ?", uid).UpdateColumn("follower_count", gorm.Expr("follow_count + ?", 1)).Error
	if err != nil {
		Logger.Error("inc user follow count error")
		return err
	}
	return nil
}

func (*UserDao) DecUserFollow(uid int64) error {
	err := db.Model(User{}).Where("id = ?", uid).UpdateColumn("follower_count", gorm.Expr("follow_count - ?", 1)).Error
	if err != nil {
		Logger.Error("dec user follow count error")
		return err
	}
	return nil
}

func (*UserDao) IncUserFollower(uid int64) error {
	err := db.Model(User{}).Where("id = ?", uid).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error
	if err != nil {
		Logger.Error("inc user follower count error")
		return err
	}
	return nil
}

func (*UserDao) DecUserFollower(uid int64) error {
	err := db.Model(User{}).Where("id = ?", uid).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	if err != nil {
		Logger.Error("dec user follower count error")
		return err
	}
	return nil
}
