package serializer

import (
	"gin_mall/conf"
	"gin_mall/repository/db/model"
)

type Product struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"` //记录商品的点击次数
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossID        int    `json:"boss_id"`
	BossName      string `json:"boss_name"`
	//BossAvatar    string `json:"boss_avatar"` // 头像
}

// BuildProduct 序列化商品
func BuildProduct(product *model.Product) *Product {
	p := &Product{
		ID:            product.ID,
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		View:          product.View(),
		CreatedAt:     product.CreatedAt.Unix(),
		Num:           product.Num,
		OnSale:        product.OnSale,
		BossID:        int(product.BossID),
		BossName:      product.BossName,
	}
	return p
}

func BuildProducts(items []*model.Product) (products []*Product) {
	for _, val := range items {
		item := BuildProduct(val)
		products = append(products, item)
	}
	return
}
