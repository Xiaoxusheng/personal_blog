package models

import "personal_blog/db"

type Blog_posts struct {
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
	_, err := db.DB.Exec("insert  into blog_post(identification,content,status,title,create_time,ip,update_time,category) value (?,?,?,?,?,?,?,?)", identification, content, status, title, create_time, ip, update_time, category)
	if err != nil {
		return err
	}
	return nil
}

func GetByTitle(title string) error {
	b := Blog_posts{}
	err := db.DB.Get(&b, "select * from blog_post wheretitle=? ", title)
	if err != nil {
		return err
	}
	return nil
}
