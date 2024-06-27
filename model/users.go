package model

import (
	"fmt"
	"image-transporter/global"
	"image-transporter/utils"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserName   string `json:"username"`
	Password   string `json:"password"`
}

func (user *Users)AddUser() error {
	fmt.Printf("\n\nUserName:\t%s\nPassword:\t%s\n\n", user.UserName, user.Password)
	if hashedPassword, err := utils.HashPassword(user.Password); err != nil {
		return err
	}else {
		user.Password = hashedPassword
		return global.DB.FirstOrCreate(&user).Error
	}
}
