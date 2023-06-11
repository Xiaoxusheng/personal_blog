package models

import (
	"fmt"
	"personal_blog/db"
	"time"
)

type Comment struct {
	ID           int       `json:"id"`
	Article_iD   string    `json:"article_id"`
	User_iD      string    `json:"user_id"`
	Parent_iD    string    `json:"parent_id"`
	Comment_id   string    `json:"comment_id"`
	Content      string    `json:"content"`
	Status       int       `json:"status"`
	Created_time time.Time `json:"created_time"`
	Updated_time time.Time `json:"updated_time"`
}

func GetComment(article_id string) []Comment {
	comment := []Comment{}
	err := db.DB.Select(&comment, "select * from comment where article_id=?", article_id)
	if err != nil {
		return nil
	}
	return comment
}

func GetCommentByArticleID(article_id, user_id string) bool {
	comment := Comment{}
	err := db.DB.Get(&comment, "select * from comment where article_id=? and user_id=? ", article_id, user_id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("comm", comment)
	return true
}

// 添加评论
func InsertComment(article_id, user_id, parent_id, content, comment_id string) error {
	_, err := db.DB.Exec("insert into comment(article_id,user_id,parent_id,content,comment_id,status,created_time,updated_time) value (?,?,?,?,?,?,?,?)", article_id, user_id, parent_id, content, comment_id, 0, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	return nil
}

// 审核
func UpdateComment(comment_id string, status int) error {
	_, err := db.DB.Exec("update comment set status=? where comment_id=?", status, comment_id)
	if err != nil {
		return err
	}
	return nil
}

// 根据comment_id判断是否存在
func GetByCommentId(comment_id string) error {
	comment := Comment{}
	err := db.DB.Get(&comment, "select * from comment where comment_id=?", comment_id)
	if err != nil {
		return err
	}
	return nil

}
