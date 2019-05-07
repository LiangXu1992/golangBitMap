package Models

import (
	"douyin/app/Constants"
	"sync"
)

type BitMap struct {
	bit    []uint64
	maxPos uint64
}

var videoBitMapMap map[int64]map[string]*BitMap
var videoCommentLock sync.Mutex

//设置bitmap
func CreateVideoBitMap(videoId string, actionId int64, AccountId uint64) bool {
	if AccountId <= 0 {
		return false
	}
	videoCommentLock.Lock()
	defer videoCommentLock.Unlock()
	//获取video对应action的bitmap
	var bitmap = getBitMap(videoId, actionId)
	//计算位置 可以写成 var idx, pos = AccountId>>6, AccountId&0x40
	var idx, pos = AccountId / 64, AccountId % 64
	//不考虑计算idx和pos是否在边界，就让64xN处在下一个idx里
	//计算bitmap是否有足够位置
	for int(idx) > len(bitmap.bit) {
		//扩展bitmap
		bitmap.bit = append(bitmap.bit, 0)
	}
	//设置具体的bit上的值
	bitmap.bit[idx] |= 0x01 << pos
	//返回结果
	return true
}

//查找账号是否可用于任务
func CheckAccountValidForJob(videoId string, actionId int64, accountId uint64) bool {
	//计算账号的idx，pos
	var idx, pos = accountId / 64, accountId % 64
	//加载历史数据
	loadVideoBitMap(videoId)
	//获取video对应action的bitmap
	var bitmap = getBitMap(videoId, actionId)
	if bitmap.bit[idx]&0x01<<pos == 1 {
		//账号已经使用过
		return false
	} else {
		//账号未使用过，标记为已使用
		bitmap.bit[idx] |= 0x01 << pos
		return true
	}
}

//加载某个视频的bitmap
func loadVideoBitMap(videoId string) {
	var tableLog TableAccountVideoActionLog
	var where = map[string]interface{}{
		"video_id": videoId,
	}
	var rows = tableLog.GetRows(where)
	for _, v := range rows {
		CreateVideoBitMap(v.VideoId, v.ActionId, v.AccountId)
	}
}

//设置video的accountId的bit
func SetVideoBit(VideoId string, actionId int64, AccountId uint64) {
	var bitmap = getBitMap(VideoId, actionId)
	var idx, pos = calcIdxPos(AccountId)
	if len(bitmap.bit) < int(idx) {
		bitmap.bit = append(bitmap.bit, make([]uint64, Constants.DefaultBitmapCap*2, Constants.DefaultBitmapCap*2)...)
	}
	//设置具体的bit上的值
	bitmap.bit[idx] |= 0x01 << pos
}

//取消video的accountId的bit
func UnSetVideoBit(accountId uint64) (uint64, uint64) {
	var idx, pos = accountId / 64, accountId % 64
	return idx, pos
}

//获取video对应action的bitmap
func getBitMap(videoId string, actionId int64) *BitMap {
	switch actionId {
	case Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE: //video like
		if _, ok := videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE]; ok {
			videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE] = make(map[string]*BitMap)
		}
		if _, ok := videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE][videoId]; ok {
			videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE][videoId] = &BitMap{make([]uint64, Constants.DefaultBitmapLen, Constants.DefaultBitmapCap), Constants.DefaultBitmapCap}
		}
		return videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE][videoId]
	case Constants.ACCOUNT_VIDEO_ACTION_ID_COMMENT: //video comment
		if _, ok := videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_COMMENT]; ok {
			videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_COMMENT] = make(map[string]*BitMap)
		}
		if _, ok := videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_COMMENT][videoId]; ok {
			videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_COMMENT][videoId] = &BitMap{make([]uint64, Constants.DefaultBitmapLen, Constants.DefaultBitmapCap), Constants.DefaultBitmapCap}
		}
		return videoBitMapMap[Constants.ACCOUNT_VIDEO_ACTION_ID_COMMENT][videoId]
	}
	return nil
}

//计算accountId在bitmap的位置
func calcIdxPos(accountId uint64) (uint64, uint64) {
	var idx, pos = accountId / 64, accountId % 64
	return idx, pos
}
