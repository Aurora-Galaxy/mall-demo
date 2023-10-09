package dao

import (
	"context"
	"gin_mall/repository/db/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDB(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// GetProductByID 通过id获取对应的product
func (dao *ProductDao) GetProductByID(uid uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id=?", uid).First(&product).Error
	return
}

// CreateProduct 创建product
func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}
