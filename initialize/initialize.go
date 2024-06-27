package initialize

import (
	"image-transporter/global"
	"image-transporter/model"
)

func Init()  {
	global.InitBaseConfig()
	global.InitDockerClient()
	global.InitSqlite()
	model.InitModel()
}