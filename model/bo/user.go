package bo

import "time"

type UserBo struct {
	Id        int64
	Account   string
	Name      string
	PassWord  string
	LoginTime time.Time
}

func (user *UserBo) CheckPassWord(password string) bool {
	return user.PassWord == password
}
