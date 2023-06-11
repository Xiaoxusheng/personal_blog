package models

import (
	"fmt"
	"personal_blog/db"
	"time"
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
	Url            string `json:"img_url"`
	Category       string `json:"category"`
}

func InsertBlogPosts(identification, content, title, create_time, ip, update_time, url string, category int, status int) error {
	_, err := db.DB.Exec("insert  into blog(identification,content,status,title,create_time,ip,update_time,category,url) value (?,?,?,?,?,?,?,?)", identification, content, status, title, create_time, ip, update_time, category, url)
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

func GetArticleList(page, number int) []Blog {
	b := []Blog{}
	err := db.DB.Select(&b, "select * from blog limit ?,?", (page-1)*number, number)
	if err != nil {
		return nil
	}
	return b
}

func GetByIdentification(identification string) bool {
	b := Blog{}
	err := db.DB.Get(&b, "select status from blog where identification=?", identification)
	if err != nil {
		return false
	}
	//如果status为2说明已经删除
	if b.Status == 2 {
		return false

	}
	return true

}

func DeleteArticle(identification string) error {
	_, err := db.DB.Exec("update blog set status=? where identification=?  ", 2, identification)
	if err != nil {
		return err
	}
	return nil

}

func UpdateArticle(identification, content, category, title string) error {
	_, err := db.DB.Exec("update blog set content=?,category=?,title=?,update_time=? where identification=?   ", content, category, title, time.Now().UnixMicro(), identification)
	if err != nil {
		return err
	}
	return nil
}
