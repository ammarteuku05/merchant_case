package migration

import "time"

type Base struct {
	Id        string     `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedBy string     `gorm:"default:null"`
	UpdatedBy string     `gorm:"default:null"`
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	Base
	FullName string
	Email    string `gorm:"unique"`
	Password string
	Outlet   []Outlet `gorm:"ForeignKey:UserId"`
}

type Outlet struct {
	Base
	OutletName string
	Picture    string
	Product    Product `gorm:"ForeignKey:OutletId"`
	UserId     string  `gorm:"index"`
}

type Product struct {
	Base
	ProductName  string
	Price        int64
	Sku          string
	Picture      string
	OutletId     string       `gorm:"index"`
	DisplayImage ImageProduct `gorm:"ForeignKey:ProductId"`
}

type ImageProduct struct {
	Base
	DisplayImage string
	ProductId    string `gorm:"index"`
}
