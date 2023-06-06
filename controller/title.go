package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_blog/models"
	"personal_blog/utility"
	"strconv"
	"time"
)

// PingExample godoc
// @Summary 添加文章接口
// @Param title formData string true "文章标题"
// @Param category formData string true "文章分类 0为技术类 1为生活类"
// @Param content formData string true "文章内容 "
// @Param token header string true "token"
// @Schemes
// @Description title category content token 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{"code":200,"msg":"添加好友成功！"}"
// @Router  /user/addtitle    [post]
func AddTitle(c *gin.Context) {
	title := c.PostForm("title")
	category := c.PostForm("category")
	content := c.PostForm("content")
	if title == "" || category == "" || content == "" {
		panic("必填参数不能为空!")
	}
	categorys, err := strconv.Atoi(category)
	err = models.GetByTitle(title)
	if err == nil {
		panic("文章已经题目存在！")
	}
	times := time.Now().UnixMicro()
	ip := c.ClientIP()
	err = models.InsertBlogPosts(utility.SetUuid(), content, title, strconv.FormatInt(times, 10), ip, title, categorys, 0)
	if err != nil {
		panic("新增文章失败！")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功！",
	})

}
