package main

import (
	"douyin/app/Models"
	"douyin/app/Schedules"
	"douyin/config"
	"douyin/orm"
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
	//创建channel
	var accountChannel = make(chan uint64, 1)
	//新建worker
	for i := 0; i < 1; i++ {
		go worker(accountChannel)
	}
	//帐号放进去channel
	for {
		//先简单完成，把数据库内的所有有效帐号写入
		var TableAccount Models.TableAccount
		var rows = TableAccount.GetValidRows()
		for _, row := range rows {
			accountChannel <- row.Id
		}
	}
}

func worker(ch chan uint64) {
	for {
		accountId := <-ch
		//获取任务
		var boolResult = Models.GetJob(accountId)
		if boolResult == false {
			time.Sleep(time.Second * 1)
			log.Println("no do job")
		} else {
			log.Println("begin job")
		}
	}
}
