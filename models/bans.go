package models

import "personal_blog/db"

type Bans struct {
	Id   int    `json:"id"`
	Ip   string `json:"ip"`
	Time string `json:"time"`
}

func GetIp(ip string) bool {
	ban := Bans{}
	err := db.DB.Get(&ban, "select * from bans where ip=?", ip)
	if err != nil {
		return false
	}
	return true
}

func InsertIpbyBans(ip, time string) bool {
	_, err := db.DB.Exec("insert bans into (ip,time )value (?,?)", ip, time)
	if err != nil {
		return false
	}
	return true
}
