package Models

import (
	"douyin/app/Constants"
	"douyin/orm"
)

type TableAccount struct {
	Id         uint64
	Mobile     string
	CreateTime int64
	UpdateTime int64
	Session    string
}

func (t *TableAccount) TableName() string {
	return "account"
}

type TTableAccountRows []TableAccount

func (t *TableAccount) GetValidRows() TTableAccountRows {
	var rows TTableAccountRows
	orm.Gorm.Where("status = ?", Constants.VALID).Find(&rows)
	return rows
}
