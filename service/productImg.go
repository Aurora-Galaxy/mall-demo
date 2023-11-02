package service

import (
	"context"
	"gin_mall/pkg/e"
	"gin_mall/repository/db/dao"
	"gin_mall/repository/db/model"
	"gin_mall/serializer"
	logging "github.com/sirupsen/logrus"
	"strconv"
)

type ProductImgService struct {
}

// ListImgPath 获取商品图片路径
func (productImgService *ProductImgService) ListImgPath(ctx context.Context, pId string) serializer.Response {
	code := e.SUCCESS
	var productImgs []*model.ProductImg
	id, _ := strconv.Atoi(pId)
	productImgDao := dao.NewProductImgDB(ctx)
	productImgs, err := productImgDao.ListProductImg(uint(id))
	if err != nil {
		code = e.ErrorDatabase
		logging.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildList(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
