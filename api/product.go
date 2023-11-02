package api

import (
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProduct 创建商品
func CreateProduct(ctx *gin.Context) {
	//gin上传多个文件
	form, _ := ctx.MultipartForm() //解析表单数据
	files := form.File["file"]     //获取前端传入的file字段所对应的内容
	var productService service.ProductService
	claims, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	err := ctx.ShouldBind(&productService)
	if err == nil {
		res := productService.Create(ctx.Request.Context(), claims.ID, files)
		ctx.JSON(e.SUCCESS, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// ListProduct 显示商品列表
func ListProduct(ctx *gin.Context) {
	var productService service.ProductService
	err := ctx.ShouldBind(&productService)
	if err == nil {
		res := productService.List(ctx.Request.Context())
		ctx.JSON(e.SUCCESS, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// SearchProduct 搜索商品
func SearchProduct(ctx *gin.Context) {
	var productService service.ProductService
	err := ctx.ShouldBind(&productService)
	if err == nil {
		res := productService.Search(ctx.Request.Context())
		ctx.JSON(e.SUCCESS, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// GetInfoProduct 获取商品详细信息
func GetInfoProduct(ctx *gin.Context) {
	var productService service.ProductService
	res := productService.GetInfo(ctx.Request.Context(), ctx.Param("id"))
	ctx.JSON(e.SUCCESS, res)
}

// GetImgProduct 获取商品图片列表
func GetImgProduct(ctx *gin.Context) {
	var productImgService service.ProductImgService
	res := productImgService.ListImgPath(ctx.Request.Context(), ctx.Param("id"))
	ctx.JSON(e.SUCCESS, res)
}
