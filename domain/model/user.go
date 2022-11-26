package model

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email,min=6,max=100"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Age       int       `json:"age" validate:"required,min=1,max=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
    DeletedAt gorm.DeletedAt `"gorm:index" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
