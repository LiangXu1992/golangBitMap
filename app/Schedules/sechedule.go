package Schedules

import (
	"douyin/app/Constants"
	"douyin/app/Models/AccountVideoActionLog"
	"douyin/app/Models/BitMap"
)

//加载数据库日志到bitMap
func loadLog2Bitmap() {
	var Tablelog AccountVideoActionLog.TableAccountVideoActionLog
	var rows = Tablelog.GetRows(Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE)
	for _, row := range rows {
		BitMap.SetBit(row.ActionId, row.VideoId, row.AccountId)
	}

}

func Start() {
	//加载数据库日志到bitMap
	loadLog2Bitmap()

}
