package model

import "gorm.io/gorm"

// Product 商品模型
type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint   `gorm:"not null"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string //打折价格
	OnSale        bool   `gorm:"default:false"` //是否正在售卖
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}
