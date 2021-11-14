package entity

import (
	"time"
)

type Outlet struct {
	Id         string    `gorm:"PrimaryKey" json:"id"`
	OutletName string    `json:"outlet_name"`
	Picture    string    `json:"picture"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Product    []Product `json:"product"`
	UserId     string    `json:"user_id"`
}

type OutletInput struct {
	OutletName string `json:"outlet_name" binding:"required"`
	Picture    string `json:"picture" binding:"required"`
	UserId     string `json:"user_id" binding:"required"`
}
