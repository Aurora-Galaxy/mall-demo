package service

import (
	"context"
	"gin_mall/conf"
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"gin_mall/repository/db/dao"
	"gin_mall/repository/db/model"
	"gin_mall/serializer"
	logging "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type EmailService struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// OperationType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `json:"operationType"  form:"operationType"`
}

type ValidEmailService struct {
}

// Send 发送邮件
func (emailService *EmailService) Send(ctx context.Context, id uint) serializer.Response {
	code := e.SUCCESS
	var notice *model.Notice
	var address string
	//生成emailToken
	token, err := utils.GenerateEmailToken(id, emailService.Email, emailService.OperationType, emailService.Password)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//获取相应的notice
	noticeDao := dao.NewNoticeDB(ctx)
	notice, err = noticeDao.GetNoticeByID(emailService.OperationType)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//拼接发送的内容
	address = conf.Config.Email.ValidEmail + token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1) //将notice中的Email替换为address，-1代表全部替换
	err = utils.SendMail([]byte(mailText), emailService.Email)
	if err != nil {
		logging.Info(err)
		code = e.ErrorSendEmail
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

func (emailService *ValidEmailService) ValidEmail(ctx context.Context, token string) serializer.Response {
	var userId uint
	var password string
	var operationType uint
	var email string
	code := e.SUCCESS
	//验证token
	if token == "" {
		code = e.InvalidParams
	} else {
		emailClaims, err := utils.ParseEmailToken(token)
		if err != nil {
			logging.Info(err)
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > emailClaims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userId = emailClaims.ID
			password = emailClaims.Password
			operationType = emailClaims.OperationType
			email = emailClaims.Email
		}
	}
	if code != e.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//执行相应的操作
	switch operationType {
	case 1:
		user.Email = email
	case 2:
		user.Email = ""
	case 3:
		err = user.SetPassword(password)
		if err != nil {
			code = e.ErrorFailEncryption
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	//更新用户信息
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//成功后返回用户信息
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
