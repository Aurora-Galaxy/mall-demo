package service

import (
	"context"
	"gin_mall/pkg/e"
	"gin_mall/repository/db/dao"
	"gin_mall/serializer"
	logging "github.com/sirupsen/logrus"
)

type CarouselService struct {
}

// ListCarousel 商品轮播图
func (*CarouselService) ListCarousel(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	carouselDao := dao.NewCarouselDB(ctx)
	carousels, err := carouselDao.GetCarousel()
	if err != nil {
		code = e.ErrorDatabase
		logging.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildList(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
