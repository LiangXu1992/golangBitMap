package Schedules

import (
	"douyin/app/Constants"
	"douyin/app/Models/Account"
	"douyin/app/Models/AccountVideoActionLog"
	"douyin/app/Models/BitMap"
)

//加载数据库日志到bitMap
func loadLog2Bitmap() {
	var TableLog AccountVideoActionLog.TableAccountVideoActionLog
	var rows = TableLog.GetRows(Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE)
	for _, row := range rows {
		BitMap.SetBit(row.ActionId, row.VideoId, row.AccountId)
	}
}

//初始化账号channel
func initAccountChannel() {
	var accountChannel = make(chan uint64)
	var TableAccount Account.TableAccount
	var rows = TableAccount.GetValidRows()
	for _, row := range rows {
		accountChannel <- row.Id
	}
}

func Start() {
	//加载数据库日志到bitMap
	loadLog2Bitmap()
	//初始化账号channel
	initAccountChannel()
}
