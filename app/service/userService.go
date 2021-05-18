package service

import (
	"fmt"
	"common_gin/app/model"
	"common_gin/app/repository"
	"common_gin/common"
)

type IUserService interface {
	UserInfoById(id int64) (user *model.User, err error)
	DeleteUserById(id int64) (err error)
	UpdateUser(user *model.User) (err error)
	CreateUser(user *model.User) (id int64, err error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService() IUserService {
	gormDb, err := common.GetGormPool("default")
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repository.NewUserRepository(gormDb)
	//初始化表格
	//_ = userRepository.InitTable()

	return &UserService{userRepository: userRepository}
}

func (u UserService) UserInfoById(id int64) (user *model.User, err error) {

	return u.userRepository.FindUserById(id)
}

func (u UserService) DeleteUserById(id int64) (err error) {
	return u.userRepository.DeleteUserById(id)
}

func (u UserService) UpdateUser(user *model.User) (err error) {
	return u.userRepository.UpdateUser(user)
}

func (u UserService) CreateUser(user *model.User) (id int64, err error) {
	return u.userRepository.CreateUser(user)
}