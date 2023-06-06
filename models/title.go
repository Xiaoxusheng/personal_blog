package models

import (
	"fmt"
	"personal_blog/db"
)

type Blog struct {
	Id             int    `json:"id"`
	Identification string `json:"identification"`
	Content        string `json:"content"`
	Status         int    `json:"status"`
	Title          string `json:"title"`
	Create_time    string `json:"create_time"`
	IP             string `json:"ip"`
	Update_time    string `json:"update_time"`
	Category       string `json:"category"`
}

func InsertBlogPosts(identification, content, title, create_time, ip, update_time string, category int, status int) error {
	_, err := db.DB.Exec("insert  into blog(identification,content,status,title,create_time,ip,update_time,category) value (?,?,?,?,?,?,?,?)", identification, content, status, title, create_time, ip, update_time, category)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetByTitle(title string) error {
	b := Blog{}
	err := db.DB.Get(&b, "select * from blog where title=? ", title)
	if err != nil {
		return err
	}
	return nil
}
