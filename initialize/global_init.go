package initialize

import (
	"gin-starter/core/config"
	"gin-starter/core/db"
	"gin-starter/core/log"
	"gin-starter/global"
)

func GlobalInit() {
	// 初始化 配置
	global.Gol_CONFIG = config.InitConfig()
	// 初始化 日志
	global.Gol_LOG = log.InitLog()
	// 初始化 数据库
	global.Gol_DB = db.InitDB()
}
