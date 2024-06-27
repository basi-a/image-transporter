package model

import "image-transporter/global"

func InitModel()  {
	autoMigrate()
}

func autoMigrate()  {
	if !global.DB.Migrator().HasTable(&Users{}){
		global.DB.AutoMigrate(&Users{})
	}
}