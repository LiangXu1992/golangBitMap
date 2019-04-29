package Schedules

import (
	"douyin/app/Constants"
	"douyin/app/Models"
)

//加载数据库日志到bitMap
func loadLog2Bitmap() {
	var TableLog Models.TableAccountVideoActionLog
	var rows = TableLog.GetRows(Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE)
	for _, row := range rows {
		Models.SetBit(row.ActionId, row.VideoId, row.AccountId)
	}
}

func Start() {
	//加载数据库日志到bitMap
	loadLog2Bitmap()
}
