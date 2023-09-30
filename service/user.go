package service

import (
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"gin_mall/repository/db/dao"
	"gin_mall/repository/db/model"
	"gin_mall/serializer"
	"golang.org/x/net/context"
)

type UserService struct {
	Nickname string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"`
}

func (userService *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.SUCCESS
	if userService.Key == "" || len(userService.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥为空或长度不足",
		}
	}
	utils.Encrypt.SetKey(userService.Key) //获取加密密钥
	userDao := dao.NewUserDao(ctx)        //新建一个用户连接
	_, exit, err := userDao.ExistOrNotByUserName(userService.UserName)
	//用户已存在
	if exit {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user = model.User{
		UserName: userService.UserName,
		NickName: userService.Nickname,
		Status:   model.Active,                       // 激活状态
		Money:    utils.Encrypt.AesEncoding("10000"), //对初始金额加密
	}
	//用户输入的密码加密
	err = user.SetPassword(userService.Password)
	if err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}