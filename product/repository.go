package product

import (
	"merchant-service/entity"
	"merchant-service/query"
	"time"

	"gorm.io/gorm"
)

type Product interface {
	CreateProduct(product entity.Product) (entity.Product, error)
	CreatedDisplayImage(displayImage entity.ImageProduct) (entity.ImageProduct, error)
	ShowAllProduct() ([]entity.Product, error)
	FindProductByID(ID string) (entity.Product, error)
	FindProductWithImageByID(ID string) (entity.Product, error)
	UpdateProductByID(ID string, input entity.UpdateProductInput) (entity.Product, error)
	DeleteProductByID(ID string) (string, error)
	FindOutletProductByID(ID string) (entity.Outlet, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateProduct(product entity.Product) (entity.Product, error) {

	qry := query.QueryCreateProduct

	err := r.db.Raw(qry,
		product.Id,
		product.ProductName,
		product.Price,
		product.Sku,
		product.Picture,
		product.CreatedAt,
		product.UpdatedAt,
		product.OutletId).Scan(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) CreatedDisplayImage(displayImage entity.ImageProduct) (entity.ImageProduct, error) {
	qry := query.QueryCreateImage

	err := r.db.Raw(qry,
		displayImage.Id,
		displayImage.DisplayImage,
		displayImage.ProductId).Scan(&displayImage).Error

	if err != nil {
		return displayImage, err
	}

	return displayImage, nil
}

func (r *repository) ShowAllProduct() ([]entity.Product, error) {
	var product []entity.Product

	qry := query.QueryFindAllProduct

	err := r.db.Raw(qry).Scan(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindProductByID(ID string) (entity.Product, error) {
	var product entity.Product

	qry := query.QueryFindProductById

	err := r.db.Raw(qry, ID).Scan(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindProductWithImageByID(ID string) (entity.Product, error) {
	var product entity.Product

	err := r.db.Where("id = ?", ID).Preload("ImageProduct").Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) UpdateProductByID(ID string, input entity.UpdateProductInput) (entity.Product, error) {

	var product entity.Product

	input.UpdatedAt = time.Now()

	qry := query.QueryUpdateProductByID
	err := r.db.Raw(qry,
		input.ProductName,
		input.Price,
		input.Sku,
		input.Picture,
		input.OutletId,
		input.UpdatedAt,
		ID).Scan(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) DeleteProductByID(ID string) (string, error) {
	product := &entity.Product{}
	qry := query.QueryDeleteProductById

	err := r.db.Raw(qry, ID).Scan(&product).Error
	if err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *repository) FindOutletProductByID(ID string) (entity.Outlet, error) {
	var outlet entity.Outlet

	err := r.db.Where("id = ?", ID).Preload("Product").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
