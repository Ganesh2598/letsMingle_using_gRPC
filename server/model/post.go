package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Description string `json:"description" binding:"required"`
	ImageUrl string `json:"imageUrl" binding:"required"`
	Email string `json:"email" binding:"required"`
	Option string `json:"option" binding:"required"`
	Username string `json:"username" binding:"required"`
	Comments []Comment `gorm:"ForeignKey:PostId"`
}