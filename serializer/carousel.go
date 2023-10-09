package serializer

import "gin_mall/repository/db/model"

type Carousel struct {
	Id        uint   `json:"id" form:"id"`
	ImagePath string `json:"imagePath" form:"imagePath"`
	ProductId uint   `json:"productId" form:"productId"`
	CreateAt  int64  `json:"createAt" form:"createAt"`
}

func BuildCarousel(carousel *model.Carousel) *Carousel {
	return &Carousel{
		Id:        carousel.ID,
		ImagePath: carousel.ImgPath,
		ProductId: carousel.ProductID,
		CreateAt:  carousel.CreatedAt.Unix(),
	}
}

func BuildCarousels(carousel []*model.Carousel) (carousels []*Carousel) {
	for _, item := range carousel {
		temp := BuildCarousel(item)
		carousels = append(carousels, temp)
	}
	return
}
