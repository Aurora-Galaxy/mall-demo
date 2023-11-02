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
	"strconv"
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

// Create 创建商品
func (productImgService *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	//multipart.FileHeader 主要用来处理文件上传
	var boss *model.User //商家
	code := e.SUCCESS
	bossDao := dao.NewUserDao(ctx)
	boss, _ = bossDao.GetUserById(uId)
	//以第一张图片作为封面图
	tmp, _ := files[0].Open() // 返回一个io.reader,用来读取文件
	path, err := utils.UploadProductToLocalStatic(tmp, uId, productImgService.Name)
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
		Name:          productImgService.Name,
		CategoryID:    uint(productImgService.CategoryID),
		Title:         productImgService.Title,
		Info:          productImgService.Info,
		ImgPath:       path,
		Price:         productImgService.Price,
		DiscountPrice: productImgService.DiscountPrice,
		Num:           productImgService.Num,
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
	//上传商品图片
	wg := new(sync.WaitGroup) //new一个等待组，确保一组任务都完成
	wg.Add(len(files))
	for index, file := range files {
		number := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB) //复用创建商品的数据库连接
		tmp, _ := file.Open()
		//为了区分商品图片，可以将number改为时间，更好标识图片
		path, err = utils.UploadProductToLocalStatic(tmp, uId, productImgService.Name+number)
		if err != nil {
			logging.Info(err)
			code = e.ErrorUploadFile
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}

// List 商品列表
func (productImgService *ProductService) List(ctx context.Context) serializer.Response {
	var products []*model.Product
	code := e.SUCCESS
	if productImgService.PageSize == 0 {
		productImgService.PageSize = 15
	}
	//商品筛选条件
	condition := make(map[string]interface{})
	if productImgService.CategoryID != 0 {
		condition["category_id"] = productImgService.CategoryID //按照商品类别筛选
	}
	productDao := dao.NewProductDB(ctx)
	count, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao2 := dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao2.ListProductByCondition(condition, productImgService.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildList(serializer.BuildProducts(products), uint(count))
}

// Search 搜索商品
func (productImgService *ProductService) Search(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	var products []*model.Product
	if productImgService.PageSize == 0 {
		productImgService.PageSize = 15
	}
	productDao := dao.NewProductDB(ctx)
	products, err := productDao.SearchProduct(productImgService.Info, productImgService.BasePage)
	if err != nil {
		code = e.ErrorDatabase
		logging.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildList(serializer.BuildProducts(products), uint(len(products)))

}

// GetInfo 获取商品信息
func (productImgService *ProductService) GetInfo(ctx context.Context, pId string) serializer.Response {
	code := e.SUCCESS
	id, _ := strconv.Atoi(pId) //将string转为int
	var product *model.Product
	productDao := dao.NewProductDB(ctx)
	product, err := productDao.GetInfo(uint(id))
	if err != nil {
		code = e.ErrorDatabase
		logging.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
	}
}
