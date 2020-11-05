package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Comment string `json:"comment" binding:"required"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	PostId uint32 `json:"postid" binding:"required"`
}