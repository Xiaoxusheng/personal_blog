package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_blog/db"
	"personal_blog/models"
	"personal_blog/utility"
)

// TODO 用户登录

// @Summary 登录接口
// @Param username formData  string true "用户名"
// @Param password formData string true "密码"
// @Schemes
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Description 用户名 密码 为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
// @Success 200 { object } "{"code": 200, "msg": "登陆成功","token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpbmRlbnRseSI6IjZhMmE0NjJjLWExMDctNDhlYS04MmU1LTc0ZTMwODMyN2U2ZiIsInVzZXJuYW1lIjoiYWRtaW4iLCJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjc4Nzg2NTM1fQ.P4dJ_f2UGhKbpiIqHxTxghRKwKIlCpF2XjryHCSnKKk"}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		panic("必填参数不能为空!")
	}
	f := models.GetUser(username, password)
	fmt.Println("f", f.Identification)
	if f.Identification == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "密码错误或账号不存在",
		})
		return
	}
	token := utility.GetToken(f.Identification)
	fmt.Println("token", token)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登陆成功！",
		"data": gin.H{
			"token": token,
		},
	})

}

// @Summary 注册接口
// @Param username formData  string true "用户名"
// @Param password formData string true "密码"
// @Param email formData string true "邮箱"
// @Schemes
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Description 用户名 密码 邮箱 为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
//
// @Success 200 { object } "{	"code": 200,	"msg": "注册成功！" 	}"
//
// @Router /user/register [post]
// TODO 用户注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	if username == "" || password == "" || email == "" {
		panic("必填参数不能为空!")
	}
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		panic(err.Error())
	}

	if f := models.GetUserByUsername(username); f {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户名已经存在！请选择其他昵称",
		})
		return
	}
	if models.GetByEmail(email) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "邮箱已经注册！",
		})
		return
	}
	err := models.InsertUser(utility.SetUuid(), username, utility.GetMa5(password), email)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功！",
	})
}

// @Summary 退出登录接口
// @Schemes
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Param Authorization header string true "token"
// @Description token 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
//
//	@Success 200 { object } "{ "code": 200, "msg": "退出登录成功！" }"
//
// @Router /user/logout [get]
func Logout(c *gin.Context) {
	identification, exists := c.Get("Identification")
	ctx := context.Background()
	if exists {
		_, err := db.Rdb.Del(ctx, identification.(string)).Result()
		if err != nil {
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出登录成功！",
	})

}
