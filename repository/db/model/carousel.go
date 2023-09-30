package model

import "gorm.io/gorm"

// Carousel 轮播图
type Carousel struct {
	gorm.Model
	ImgPath   string // 图片路径
	ProductID uint   `gorm:"not null"` // 产品id
}
