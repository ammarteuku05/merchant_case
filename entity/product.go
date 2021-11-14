package entity

import (
	"time"
)

type Product struct {
	Id           string         `gorm:"PrimaryKey" json:"id"`
	ProductName  string         `json:"product_name"`
	Price        int            `json:"price"`
	Sku          string         `json:"sku"`
	Picture      string         `json:"picture"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	OutletId     string         `json:"outlet_id"`
	ImageProduct []ImageProduct `json:"imageproduct"`
}

type ImageProduct struct {
	Id           string `json:"id"`
	DisplayImage string `json:"display_image"`
	ProductId    string `json:"product_id"`
}

type ProductInput struct {
	ProductName string `json:"product_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Sku         string `json:"sku" binding:"required"`
	Picture     string `json:"picture" binding:"required"`
	OutletId    string `json:"outlet_id" binding:"required"`
}

type UpdateProductInput struct {
	ProductName string    `json:"product_name"`
	Price       int64     `json:"price" `
	Sku         string    `json:"sku"`
	Picture     string    `json:"picture"`
	OutletId    string    `json:"outlet_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
