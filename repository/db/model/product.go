package model

import (
	"gin_mall/repository/cache"
	"gorm.io/gorm"
	"strconv"
)

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
	// BossAvatar    string
}

// View 记录用户点击商品数量
func (product *Product) View() uint64 {
	// 根据key从Redis中取出点击数
	countStr, _ := cache.RedisClient.Get(cache.ProductKeyView(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	//增加商品点击数
	cache.RedisClient.Incr(cache.ProductKeyView(product.ID))
}
