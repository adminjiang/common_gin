package controller

import (
	"github.com/gin-gonic/gin"
	"common_gin/app/service"
	"common_gin/lib"
	"strconv"
)

type UserController struct {
	userService service.IUserService
}

func UserRegister(group *gin.RouterGroup) {
	//注册UserService
	userService := service.NewUserService()
	admin := &UserController{userService: userService}
	group.GET("/userInfo", admin.UserInfo)
}

func (u UserController) UserInfo(c *gin.Context) {
	id := c.Query("id")
	int64Id, _ := strconv.ParseInt(id, 10, 64)
	useInfo, err := u.userService.UserInfoById(int64Id)
	if err != nil {
		lib.ResponseError(c,400,err)
		return
	}
	lib.ResponseSuccess(c,useInfo)
}
