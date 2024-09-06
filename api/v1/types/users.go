package types

import (
	"time"
)

type UserBody struct {
	Username string    `json:"username" validate:"required,min=3,max=20"`
	Password string    `json:"password" validate:"required,min=6"`
	Email    string    `json:"email" validate:"required,email"`
	Name     string    `json:"name" validate:"required"`
	Gender   string    `json:"gender" validate:"required,oneof=male female other"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}