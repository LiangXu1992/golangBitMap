package Models

import (
	"douyin/app/Constants"
	"sync"
)

type BitMap struct {
	bit    []uint64
	maxPos uint64
}

var VideoCommentMap map[string]*BitMap = make(map[string]*BitMap)
var VideoLikeMap map[string]*BitMap = make(map[string]*BitMap)
var videoCommentLock sync.Mutex

//设置bitmap
func SetBit(VideoId string, actionId int64, AccountId uint64) bool {
	if AccountId <= 0 {
		return false
	}
	videoCommentLock.Lock()
	defer videoCommentLock.Unlock()
	switch actionId {
	case Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE:
		var _, ok = VideoLikeMap[VideoId]
		if !ok {
			//创建视频like的bitmap
			VideoLikeMap[VideoId] = &BitMap{make([]uint64, Constants.DefaultBitmapLen, Constants.DefaultBitmapCap), Constants.DefaultBitmapCap}
		}
		//计算位置 可以写成 var idx, pos = AccountId>>6, AccountId&0x40
		var idx, pos = AccountId / 64, AccountId % 64
		//不考虑计算idx和pos是否在边界，就让64xN处在下一个idx里
		//计算bitmap是否有足够位置
		for int(idx) > len(VideoLikeMap[VideoId].bit) {
			//扩展bitmap
			VideoLikeMap[VideoId].bit = append(VideoLikeMap[VideoId].bit, 0)
		}
		//设置具体的bit上的值
		VideoLikeMap[VideoId].bit[idx] |= 0x01 << pos
		//返回结果
		return true
	case Constants.ACCOUNT_VIDEO_ACTION_ID_COMMENT:
		var _, ok = VideoCommentMap[VideoId]
		if !ok {
			//创建视频like的bitmap
			VideoCommentMap[VideoId] = &BitMap{make([]uint64, Constants.DefaultBitmapLen, Constants.DefaultBitmapCap), Constants.DefaultBitmapCap}
		}
		//计算位置 可以写成 var idx, pos = AccountId>>6, AccountId&0x40
		var idx, pos = AccountId / 64, AccountId % 64
		//不考虑计算idx和pos是否在边界，就让64xN处在下一个idx里
		//计算bitmap是否有足够位置
		for int(idx) > len(VideoCommentMap[VideoId].bit) {
			//扩展bitmap
			VideoCommentMap[VideoId].bit = append(VideoCommentMap[VideoId].bit, 0)
		}
		//设置具体的bit上的值
		VideoCommentMap[VideoId].bit[idx] |= 0x01 << pos
		//返回结果
		return false
	default:
		return false
	}
}

//查找账号是否可用于任务
func CheckAccountValidForJob(videoId string, actionId int64, accountId uint64) bool {
	//计算账号的idx，pos
	var idx, pos = accountId / 64, accountId % 64
	switch actionId {
	case Constants.ACCOUNT_VIDEO_ACTION_ID_LIKE: //video like
		//获取bit上的值，进行&计算
		var _, ok = VideoLikeMap[videoId]
		if !ok {
			//创建视频like的bitmap
			VideoLikeMap[videoId] = &BitMap{make([]uint64, Constants.DefaultBitmapLen, Constants.DefaultBitmapCap), Constants.DefaultBitmapCap}
		}
		if VideoLikeMap[videoId].bit[idx]&0x01<<pos == 1 {
			//账号已经使用过
			return false
		} else {
			//账号未使用过，标记为已使用
			VideoLikeMap[videoId].bit[idx] |= 0x01 << pos
			return true
		}
	case Constants.ACCOUNT_VIDEO_ACTION_ID_COMMENT: //video comment
		var _, ok = VideoCommentMap[videoId]
		if !ok {
			//创建视频like的bitmap
			VideoCommentMap[videoId] = &BitMap{make([]uint64, Constants.DefaultBitmapLen, Constants.DefaultBitmapCap), Constants.DefaultBitmapCap}
		}
		//获取bit上的值，进行&计算
		if VideoCommentMap[videoId].bit[idx]&0x01<<pos == 1 {
			//账号已经使用过
			return false
		} else {
			//账号未使用过，标记为已使用
			VideoCommentMap[videoId].bit[idx] |= 0x01 << pos
			return true
		}
	}
	//返回结果
	return false
}
