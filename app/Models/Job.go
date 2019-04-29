package Models

import (
	"douyin/app/Constants"
	"douyin/orm"
	"errors"
	"log"
	"time"
)

type TableJob struct {
	Id           uint64
	Title        string
	VideoId      string
	ActionId     int64
	CreateTime   int64
	Status       int64
	IsFinish     int64
	TotalNumber  int64
	ActualNumber int64
}

func (t *TableJob) TableName() string {
	return "job"
}

type JobStruct struct {
	JobMap map[uint64]*TableJob
}

//执行任务
func (job *JobStruct) ExecJob(ch chan uint64) bool {
	var accountId = <-ch
	//获取一个任务，任务有很多类型，但都是以帐号为主体去执行
	var tJob, err = job.getJob(accountId)
	if err != nil {
		return false
	}
	//正式执行任务
	log.Println(tJob)
	time.Sleep(time.Second * 100)
	return true
}

//获取一个任务
func (job *JobStruct) getJob(accountId uint64) (*TableJob, error) {
	for _, v := range job.JobMap {
		//额外补多一百次，以防重复执行太多
		if v.ActualNumber < v.TotalNumber+100 {
			//判断帐号是否可用于任务
			if CheckAccountValidForJob(v.VideoId, v.ActionId, accountId) == true {
				//setBit
				if SetBit(v.VideoId, v.ActionId, accountId) == true {
					//incr actual_number
					v.ActualNumber += 1
					//返回结果
					return v, nil
				}
			}
			return v, nil
		}
	}
	return nil, errors.New("no job")
}

//读取mysql的未完成的job
func (job *JobStruct) LoadHistoryJob() {
	var jobSlice []TableJob
	orm.Gorm.Where("is_finish = ?", Constants.INVALID).Find(&jobSlice)
	for _, v := range jobSlice {
		job.JobMap[v.Id] = &v
	}
}
