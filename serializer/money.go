package serializer

import (
	"gin_mall/pkg/utils"
	"gin_mall/repository/db/model"
)

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(user *model.User, key string) Money {
	utils.Encrypt.SetKey(key)
	return Money{
		UserId:    user.ID,
		UserName:  user.UserName,
		UserMoney: utils.Encrypt.AesDecoding(user.Money),
	}
}
