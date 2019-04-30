package main

import (
	"douyin/app/Models"
	"douyin/app/Schedules"
	"douyin/config"
	"douyin/orm"
	"douyin/routes"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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
	var TableAccount Models.TableAccount
	//创建accountChannel
	var tmpRow = orm.Gorm.Select("max(id) max_id").Row()
	var accountChannelLen = 0
	_ = tmpRow.Scan(&accountChannelLen)
	//防止channel阻塞
	var accountChannel = make(chan uint64, accountChannelLen*2)
	//新建jobStruct
	var job = Models.JobStruct{
		JobMap: make(map[uint64]*Models.TableJob),
	}
	job.LoadHistoryJob()
	//新建worker
	for i := 0; i < 10; i++ {
		go worker(&job, accountChannel)
	}
	//帐号放进去channel
	var rows = TableAccount.GetValidRows()
	for _, row := range rows {
		accountChannel <- row.Id
	}

	//api服务
	//初始化路由
	router := gin.New()
	routes.Start(router)
	//启动端口服务
	_ = router.Run(":6789") // listen and serve on 0.0.0.0:8080
}

func worker(job *Models.JobStruct, ch chan uint64) {
	for {
		//执行任务
		var boolResult = job.ExecJob(ch)
		if boolResult == false {
			time.Sleep(time.Second * 1)
			log.Println("no do job")
		} else {
			log.Println("begin job")
		}
	}
}
