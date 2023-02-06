package dao

import "time"

type UserDao struct {
	Account   string
	Password  string
	LoginTime time.Time
	BaseDao
}

func (UserDao) TableName() string {
	return "user"
}
