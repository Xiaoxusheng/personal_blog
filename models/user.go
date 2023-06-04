package models

import (
	"personal_blog/db"
	"time"
)

type User struct {
	Id                int    `json:"id" `
	Identification    string `json:"identification" `
	Username          string `json:"username" form:"username"  binding:"required,min=3,max=10" form:"username"`
	Password          string `json:"password" form:"password" binding:"required,min=5,max=10"  form:"password" `
	Email             string `json:"email" form:"email" binding:"required,email"`
	Status            int    `json:"status" `
	Registration_time string `json:"registration_Time" form:"registration_Time"`
}

func GetUserByUsername(username string) bool {
	user := User{}
	err := db.DB.Get(&user, "select * from user where username=?", username)
	if err != nil {
		return false
	}
	return true
}

func GetUser(username, password string) User {
	user := User{}
	err := db.DB.Get(&user, "select * from user where username=? and password=?", username, password)
	if err != nil {
	}
	return user

}

func GetByEmail(email string) bool {
	user := User{}
	err := db.DB.Get(&user, "select * from user where email=?", email)
	if err != nil {
		return false
	}
	return true
}

func InsertUser(identification, username, password, email string) error {
	_, err := db.DB.Exec("insert into user(identification,username,password,email,registration_Time,status) value (?,?,?,?,?,?)", identification, username, password, email, time.Now().Format("2006-01-02 15:04:05"), 0)
	if err != nil {
		return err
	}
	return nil
}
