package global

import (
	"fmt"
	"image-transporter/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
func InitSqlite()  {
	dsn := fmt.Sprintf("%s/app.db", Config.BaseConfig.DbDir)
	var err error
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		utils.FatalfErr("open sqlite file faild.", err)
	}
	db, _ := DB.DB()
	db.SetMaxOpenConns(1024)
	db.SetConnMaxIdleTime(50)
}