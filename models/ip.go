package models

import (
	"personal_blog/db"
	"time"
)

type Ip struct {
	Id       int    `json:"id"`
	Ip       string `json:"ip"`
	Time     string `json:"time"`
	UseAgent string `json:"use_agent"`
}

func GetIpNumber() int {
	var num int
	//获取这个小时的时间
	times := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Now().Location()).Unix()
	err := db.DB.Get(&num, "select  count(*) from  ip where time >?", times)
	if err != nil {
		return 0
	}
	return num
}

func InsertIP(ip *Ip) error {
	_, err := db.DB.Exec("insert  into ip(ip,time,use_agent)value (?,?,?)", ip.Ip, ip.Time, ip.UseAgent)
	if err != nil {
		return err
	}
	return nil
}
