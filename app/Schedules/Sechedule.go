package Schedules

import (
	"douyin/app/Models"
)

//加载数据库日志到bitMap
func loadLog2Bitmap() {
	var TableLog Models.TableAccountVideoActionLog
	var rows = TableLog.GetRows(map[string]interface{}{})
	for _, row := range rows {
		Models.CreateVideoBitMap(row.VideoId, row.ActionId, row.AccountId)
	}
}

func Start() {
	//加载数据库日志到bitMap
	loadLog2Bitmap()
}
