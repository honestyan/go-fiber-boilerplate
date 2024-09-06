package types

import (
	"time"
)

type ProductBody struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Price       int       `json:"price" validate:"required"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
}