package dao

import "time"

type BaseDao struct {
	Id         int64
	CreateTime time.Time
	UpdateTime time.Time
}
