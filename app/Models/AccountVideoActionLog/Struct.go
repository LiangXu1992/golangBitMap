package AccountVideoActionLog

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

func (t *TableAccountVideoActionLog) GetRows(ActionId int64) TableAccountVideoActionLogRows {
	var rows TableAccountVideoActionLogRows
	orm.Gorm.Where("action_id = ?", ActionId).Find(&rows)
	return rows
}