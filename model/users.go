package model

import (
	"image-transporter/global"
	"image-transporter/utils"
	"time"

	"gorm.io/gorm"
)

const deletedUserRetentionMonths int = 6

type Users struct {
	gorm.Model
	UserName   string `json:"username" gorm:"not null;uniqueIndex"`
	Password   string `json:"password" gorm:"not null"`
	Permission string `json:"permission" gorm:"default:user;not null"`
}

func (user *Users) AddUser() error {

	if hashedPassword, err := utils.HashPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = hashedPassword
		return global.DB.FirstOrCreate(&user).Error
	}
}

// Delete users who have been soft deleted for more than 6 months
func RemoveDelatedUsers() error {
	tx := global.DB.Begin()
	if err := tx.Delete(&Users{}).Where("delete_at < ?", time.Now().AddDate(0, -deletedUserRetentionMonths, 0)).Unscoped().Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func UsersList(limit int, permission string) (users []Users, err error) {
	if permission != "" {
		if err := global.DB.Find(&users).Where("permission = ?", permission).Limit(limit).Error; err != nil {
			return nil, err
		}
	} else {
		if err := global.DB.Find(&users).Limit(limit).Error; err != nil {
			return nil, err
		}
	}

	return
}
