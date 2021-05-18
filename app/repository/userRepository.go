package repository

import (
	"github.com/jinzhu/gorm"
	"common_gin/app/model"
)

type IUserRepository interface {
	InitTable() error
	FindUserById(id int64) (user *model.User, err error)
	CreateUser(user *model.User) (id int64, err error)
	DeleteUserById(id int64) (err error)
	UpdateUser(user *model.User) (err error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) InitTable() error {
	return u.db.CreateTable(&model.User{}).Error
}

func (u UserRepository) FindUserById(id int64) (user *model.User, err error) {
	user = &model.User{}
	err = u.db.First(user, id).Error
	return user, err
}

func (u UserRepository) CreateUser(user *model.User) (id int64, err error) {
	return user.ID, u.db.Create(user).Error
}

func (u UserRepository) DeleteUserById(id int64) (err error) {
	return u.db.Where("id = ?", id).Delete(&model.User{}).Error
}

func (u UserRepository) UpdateUser(user *model.User) (err error) {
	return u.db.Model(user).Update(user).Error
}