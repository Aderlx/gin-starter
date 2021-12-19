package main

import (
	"fmt"
	"gin-starter/apps/demoapp"
	"gin-starter/global"
	"gin-starter/initialize"
	"gin-starter/router"
)

func main() {
	// 初始化全局参数
	initialize.GlobalInit()

	// 迁移数据库
	initialize.AutoMigrate()

	// 获取端口
	var port = global.Gol_CONFIG.GetString("httpServer.port")

	// 加载路由
	router.Include(demoapp.Routers)
	// 初始化路由
	r := router.Init()
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("startup service failed, err:%s\n", err.Error())
	}
}
