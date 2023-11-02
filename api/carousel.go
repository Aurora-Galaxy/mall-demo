package api

import (
	"gin_mall/pkg/e"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListCarousel 轮播图
func ListCarousel(ctx *gin.Context) {
	var carouselService service.CarouselService
	err := ctx.ShouldBind(&carouselService)
	if err == nil {
		res := carouselService.ListCarousel(ctx.Request.Context())
		ctx.JSON(e.SUCCESS, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
