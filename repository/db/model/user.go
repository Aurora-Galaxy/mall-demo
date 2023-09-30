package model

import (
	"gin_mall/conf"
	"github.com/CocaineCong/secret"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string //加密后的密码
	NickName       string //昵称
	Status         string //用户状态，是否被封禁
	//Avatar         string `gorm:"size:1000"` //用户头像
	Money     string
	Relations []User `gorm:"many2many:relation;"` //好友关系
}

const (
	PassWordCost        = 12       // 密码加密难度
	Active       string = "active" // 激活用户
)

// SetPassword 设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarURL 头像地址
//func (u *User) AvatarURL() string {
//	if conf.Config.System.UploadModel == consts.UploadModelOss {
//		return u.Avatar
//	}
//	pConfig := conf.Config.PhotoPath
//	return pConfig.PhotoHost + conf.Config.System.HttpPort + pConfig.AvatarPath + u.Avatar
//}

// EncryptMoney 加密金额
func (u *User) EncryptMoney(key string) (money string, err error) {
	aesObj, err := secret.NewAesEncrypt(conf.Config.EncryptSecret.MoneySecret, key, "", secret.AesEncrypt128, secret.AesModeTypeCBC)
	if err != nil {
		return
	}
	money = aesObj.SecretEncrypt(u.Money)

	return
}

// DecryptMoney 解密金额
func (u *User) DecryptMoney(key string) (money float64, err error) {
	aesObj, err := secret.NewAesEncrypt(conf.Config.EncryptSecret.MoneySecret, key, "", secret.AesEncrypt128, secret.AesModeTypeCBC)
	if err != nil {
		return
	}

	money = cast.ToFloat64(aesObj.SecretDecrypt(u.Money))
	return
}
