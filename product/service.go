package product

import (
	"errors"
	"fmt"
	"merchant-service/entity"
	"merchant-service/utils/formatter"
	"time"

	"github.com/gofrs/uuid"
)

type ProductService interface {
	CreateProduct(product entity.ProductInput) (formatter.ProductFormat, error)
	CreateDisplayImageProduct(pathFile string, inputproductID string) (entity.ImageProduct, error)
	ShowAllProduct() ([]formatter.ProductFormat, error)
	FindProductByID(productID string) (entity.Product, error)
	UpdateProductByID(productID string, input entity.UpdateProductInput) (formatter.ProductFormat, error)
	DeleteProductByID(productID string) (interface{}, error)
	FindOutletUserByID(outletID string) (entity.Outlet, error)
}

type productservice struct {
	repo Product
}

func NewProductService(repo Product) *productservice {
	return &productservice{repo}
}

func (s *productservice) CreateProduct(product entity.ProductInput) (formatter.ProductFormat, error) {

	productuuid, err := uuid.NewV4()

	if err != nil {
		return formatter.ProductFormat{}, err
	}

	var newProduct = entity.Product{
		Id:          productuuid.String(),
		ProductName: product.ProductName,
		Price:       product.Price,
		Sku:         product.Sku,
		Picture:     product.Picture,
		OutletId:    product.OutletId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createProduct, err := s.repo.CreateProduct(newProduct)

	formatProduct := formatter.FormatProduct(createProduct)

	if err != nil {
		return formatProduct, err
	}

	return formatProduct, nil
}

func (s *productservice) CreateDisplayImageProduct(pathFile string, inputproductID string) (entity.ImageProduct, error) {

	imageuuid, err := uuid.NewV4()

	if err != nil {
		return entity.ImageProduct{}, err
	}

	newDisplayImage := entity.ImageProduct{
		Id:           imageuuid.String(),
		DisplayImage: pathFile,
		ProductId:    inputproductID,
	}

	displayImage, err := s.repo.CreatedDisplayImage(newDisplayImage)

	if err != nil {
		return displayImage, err
	}

	return displayImage, nil
}

func (s *productservice) ShowAllProduct() ([]formatter.ProductFormat, error) {
	product, err := s.repo.ShowAllProduct()

	var formatuserProduct []formatter.ProductFormat

	for _, products := range product {
		formatProduct := formatter.FormatProduct(products)
		formatuserProduct = append(formatuserProduct, formatProduct)

	}
	if err != nil {
		return formatuserProduct, err
	}

	return formatuserProduct, nil
}

func (s *productservice) FindProductByID(productID string) (entity.Product, error) {
	product, err := s.repo.FindProductWithImageByID(productID)

	if err != nil {
		return entity.Product{}, err
	}

	if len(product.Id) == 0 {
		newError := fmt.Sprintf("product id %s not found", productID)
		return entity.Product{}, errors.New(newError)
	}

	return product, nil
}

func (s *productservice) UpdateProductByID(productID string, input entity.UpdateProductInput) (formatter.ProductFormat, error) {
	var dataUpdate = map[string]interface{}{}

	product, err := s.repo.FindProductByID(productID)

	if err != nil {
		return formatter.ProductFormat{}, err
	}

	if len(product.Id) == 0 {
		newError := fmt.Sprintf("product id %s not found", productID)
		return formatter.ProductFormat{}, errors.New(newError)
	}

	if input.ProductName != "" || len(input.ProductName) != 0 {
		dataUpdate["product_name"] = input.ProductName
	}
	if input.Price != 0 {
		dataUpdate["price"] = input.Price
	}
	if input.Sku != "" || len(input.Sku) != 0 {
		dataUpdate["sku"] = input.Sku
	}

	if input.Picture != "" || len(input.Picture) != 0 {
		dataUpdate["picture"] = input.Picture
	}
	if input.OutletId != "" || len(input.OutletId) != 0 {
		dataUpdate["OutletID"] = input.OutletId
	}
	dataUpdate["updated_at"] = time.Now()

	productUpdated, err := s.repo.UpdateProductByID(productID, input)

	if err != nil {
		return formatter.ProductFormat{}, err
	}

	formatProduct := formatter.FormatProduct(productUpdated)

	return formatProduct, nil
}

func (s *productservice) DeleteProductByID(productID string) (interface{}, error) {

	product, err := s.repo.FindProductByID(productID)

	if err != nil {
		return nil, err
	}

	if len(product.Id) == 0 {
		newError := fmt.Sprintf("Product id %s not found", productID)
		return nil, errors.New(newError)
	}

	status, err := s.repo.DeleteProductByID(productID)

	if err != nil {
		return nil, err
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete Product ID : %s", productID)

	formatDelete := formatter.FormatDeleteProduct(msg)

	return formatDelete, nil
}

func (s *productservice) FindOutletUserByID(outletID string) (entity.Outlet, error) {
	outlet, err := s.repo.FindOutletProductByID(outletID)

	if err != nil {
		return outlet, err
	}

	if len(outlet.Id) == 0 {
		newError := fmt.Sprintf("Outlet id %s not found", outletID)
		return outlet, errors.New(newError)
	}

	return outlet, nil
}
