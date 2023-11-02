package serializer

import (
	"gin_mall/conf"
	"gin_mall/repository/db/model"
)

type ProductImg struct {
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(productImg *model.ProductImg) *ProductImg {
	return &ProductImg{
		ProductID: productImg.ProductID,
		ImgPath: conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath +
			productImg.ImgPath,
	}
}

func BuildProductImgs(productImgs []*model.ProductImg) (ProductImgs []*ProductImg) {
	for _, val := range productImgs {
		item := BuildProductImg(val)
		ProductImgs = append(ProductImgs, item)
	}
	return
}
