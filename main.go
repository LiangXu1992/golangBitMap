package main

import (
	"douyin/app/Schedules"
	"douyin/config"
	"douyin/orm"
	"douyin/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	//定义日志格式
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	//读取配置文件
	config.Config.Start("config/config.develop.yml")
	//开启orm
	orm.Start()
	//初始化定时任务
	Schedules.Start()
	//初始化路由
	router := gin.New()
	routes.Start(router)
	//启动端口服务
	_ = router.Run(":6776") // listen and serve on 0.0.0.0:8080
}
