package service

import (
	"context"
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"gin_mall/repository/db/dao"
	"gin_mall/repository/db/model"
	"gin_mall/serializer"
	logging "github.com/sirupsen/logrus"
	"mime/multipart"
	"sync"
)

type ProductService struct {
	ID             uint   `form:"id" json:"id"`
	Name           string `form:"name" json:"name"`
	CategoryID     int    `form:"category_id" json:"category_id"`
	Title          string `form:"title" json:"title" `
	Info           string `form:"info" json:"info" `
	ImgPath        string `form:"img_path" json:"img_path"`
	Price          string `form:"price" json:"price"`
	DiscountPrice  string `form:"discount_price" json:"discount_price"`
	OnSale         bool   `form:"on_sale" json:"on_sale"`
	Num            int    `form:"num" json:"num"`
	model.BasePage        //分页
}

func (productService *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	//multipart.FileHeader 主要用来处理文件上传
	var boss *model.User //商家
	code := e.SUCCESS
	bossDao := dao.NewUserDao(ctx)
	boss, _ = bossDao.GetUserById(uId)
	//以第一张图片作为封面图
	tmp, _ := files[0].Open() // 返回一个io.reader,用来读取文件
	path, err := utils.UploadProductToLocalStatic(tmp, uId, productService.Name)
	if err != nil {
		logging.Info(err)
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          productService.Name,
		CategoryID:    uint(productService.CategoryID),
		Title:         productService.Title,
		Info:          productService.Info,
		ImgPath:       path,
		Price:         productService.Price,
		DiscountPrice: productService.DiscountPrice,
		Num:           productService.Num,
		OnSale:        true,
		BossID:        uId,
		BossName:      boss.UserName,
		//BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDB(ctx)
	//创建商品
	err = productDao.CreateProduct(product)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup) //new一个等待组，确保一组任务都完成
	wg.Add(len(files))
	for index, val := range files {

	}
}
