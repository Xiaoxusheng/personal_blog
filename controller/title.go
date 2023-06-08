package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_blog/models"
	"personal_blog/utility"
	"strconv"
	"time"
)

// PingExample godoc
// @Summary 添加文章接口
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Param title formData string true "文章标题"
// @Param category formData string true "文章分类 0为技术类 1为生活类"
// @Param content formData string true "文章内容 "
// @Param Authorization header string true "token"
// @Schemes
// @Description title category只能为0 或1 content token 为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
// @Success 200 {string}  "{"code":200,"msg":"添加成功！"}"
// @Router  /user/addarticle    [post]
func AddArticle(c *gin.Context) {
	title := c.PostForm("title")
	category := c.PostForm("category")
	content := c.PostForm("content")
	fmt.Println("title", title)
	if title == "" || category == "" || content == "" {
		panic("必填参数不能为空!")
	}
	if !utility.Contains(utility.List, category) {
		panic("category格式不对！")
	}
	fmt.Println(title, content, category)
	categorys, err := strconv.Atoi(category)
	err = models.GetByTitle(title)
	if err == nil {
		panic("文章已经题目存在！")
	}
	times := time.Now().UnixMicro()
	ip := c.ClientIP()
	err = models.InsertBlogPosts(utility.SetUuid(), content, title, strconv.FormatInt(times, 10), ip, strconv.FormatInt(times, 10), categorys, 0)
	if err != nil {
		panic("新增文章失败！")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功！",
	})

}

// PingExample godoc
// @Summary 获取文章接口
// @Param page query string false "页数,不填默认为1"
// @Param number query string false "每页大小，默认为20"
// @Param Authorization header string true "token"
// @Schemes
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Description   token 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{"code": 200,	"data": {"length": 4, "list": [	{	"id": 1, "identification": "1068dcc0-4fd3-461d-9f4b-0100879457eb",	"content": "1",	"status": 0,	"title": "1",	"create_time": "1686055918562468",	"ip": "127.0.0.1", "update_time": "1", "category": "1" },]},"msg": "获取数据成功！"}"
// @Router  /user/articlelist    [get]
func GetArticle(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	number := c.DefaultQuery("number", "20")
	if number == "" {
		panic("必填参数不能为空！")
	}
	p, err := strconv.Atoi(page)
	num, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	list := models.GetArticleList(p, num)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功！",
		"data": gin.H{
			"list":   list,
			"length": len(list),
		},
	})

}

// PingExample godoc
// @Summary 删除文章接口
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Param identification query string true "文章唯一标识"
// @Param Authorization header string true "token"
// @Schemes
// @Description identification token为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{ "code": 1, "err": "删除文章不存在或者已经删除！" }"
// @Router  /user/deletearticle    [delete]
func DeleteArticle(c *gin.Context) {
	identification := c.Query("identification")
	if identification == "" {
		panic("必填参数不能为空！")
	}
	if f := models.GetByIdentification(identification); !f {
		panic("删除文章不存在或者已经删除！")
	}

	err := models.DeleteArticle(identification)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功！",
	})

}

// PingExample godoc
// @Summary 更新文章接口
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Param identification query string true "文章唯一标识"
// @Param content query string true "文章内容"
// @Param category query string true "文章分类"
// @Param title query string true "文章标题"
// @Param Authorization header string true "token"
// @Schemes
// @Description identification content category title 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{ "code": 200, "msg": "更新成功！" }"
// @Router  /user/updatearticle    [get]
func UpdateArticle(c *gin.Context) {
	identification := c.Query("identification")
	content := c.Query("content")
	category := c.Query("category")
	title := c.Query("title")

	if identification == "" || content == "" || category == "" || title == "" {
		panic("必填参数不能为空！")
	}
	if !utility.Contains(utility.List, category) {
		panic("category格式不对！")
	}
	if f := models.GetByIdentification(identification); !f {
		panic("文章不存在！")
	}
	err := models.GetByTitle(title)
	if err == nil {
		panic("title已经存在！")
	}
	if err := models.UpdateArticle(identification, content, category, title); err != nil {
		panic("更新失败！" + err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功！",
	})
}
