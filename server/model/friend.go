package model

import (
	"github.com/jinzhu/gorm"
)

type Friend struct {
	gorm.Model
	FriendMail string `json:"friendmail" binding:"required"`
	FriendName string `json:"friendname" binding:"required"`
	FriendImage string `json:"friendimage" binding:"required"`
	Mymail string `json:"mymail" binding:"required"`
	MyimageUrl string `json:"myimageUrl" binding:"required"`
	Myusername string `json:"myusername" binding:"required"`
	Status string `json:"status" binding:"required"`
}

