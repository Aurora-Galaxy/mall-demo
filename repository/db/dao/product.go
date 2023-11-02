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

// CountProductByCondition 根据条件统计商品数量
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// ListProductByCondition 展示商品列表
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	//offset根据页码和每页记录数计算偏移量，以获取正确的分页数据，limit限制每页显示的数量
	err = dao.DB.Where(condition).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return
}

// SearchProduct 搜索商品，并获取其数量
func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, err error) {
	//%[info的值]% 表示一个包含指定 info 值的任意字符序列。% 在字符串中的位置是灵活的，表示可以匹配任何字符，包括零个字符
	err = dao.DB.Model(&model.Product{}).Where("info LIKE ? OR title LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return
}

// GetInfo 根据id获取商品信息
func (dao *ProductDao) GetInfo(pId uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id = ?", pId).Find(&product).Error
	return
}
