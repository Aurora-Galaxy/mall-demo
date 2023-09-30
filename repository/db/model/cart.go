package model

import "gorm.io/gorm"

// Cart 购物车模型
type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint `gorm:"not null"`
	BossID    uint // 商家id
	Num       uint
	MaxNum    uint // 最大数量
	Check     bool //判断是否支付
}
