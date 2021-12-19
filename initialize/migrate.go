package initialize

import (
	"gin-starter/apps/demoapp"
	"gin-starter/global"
)

func AutoMigrate() {

	var models = []interface{}{&demoapp.DemoModel{}}

	global.Gol_DB.AutoMigrate(models...)

}
