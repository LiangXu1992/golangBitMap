package Models

import (
	"douyin/orm"
)

type TableAccountVideoActionLog struct {
	Id         uint64
	AccountId  uint64
	VideoId    string
	ActionId   int64
	CreateTime int64
}

type TableAccountVideoActionLogRows []TableAccountVideoActionLog

func (t *TableAccountVideoActionLog) TableName() string {
	return "account_video_action_log"
}

func (t *TableAccountVideoActionLog) GetRows(where map[string]interface{}) TableAccountVideoActionLogRows {
	var rows TableAccountVideoActionLogRows
	orm.Gorm.Where(where).Find(&rows)
	return rows
}
