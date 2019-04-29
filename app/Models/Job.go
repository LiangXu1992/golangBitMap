package Models

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

func GetJob(accountId uint64) bool {
	//获取一个任务
	var jobId = "1"
	var actionId int64 = 1
	//匹配帐号与任务是否可用
	return CheckAccountValidForJob(jobId, actionId, accountId)
}
