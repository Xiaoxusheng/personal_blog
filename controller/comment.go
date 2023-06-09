package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_blog/models"
	"personal_blog/utility"
)

// PingExample godoc
// @Summary  获取评论接口
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Param article_id query string true "文章唯一标识"
// @Param Authorization header string true "token"
// @Schemes
// @Description article_id  token  为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{ "code": 200, "msg": "更新成功！" }"
// @Router  /user/commentlist    [get]
func GetCommentList(c *gin.Context) {
	article_id := c.Query("article_id")
	if article_id == "" {
		panic("必填参数不能为空！")
	}
	if f := models.GetByIdentification(article_id); !f {
		panic("文章不存在！")
	}
	list := models.GetComment(article_id)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功！",
		"data": gin.H{
			"list": list,
		},
	})

}

// PingExample godoc
// @Summary 发布评论接口
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Param article_id formData string true "文章唯一标识"
// @Param parent_id formData string false "父级评论"
// @Param content formData string true "评论内容"
// @Param Authorization header string true "token"
// @Schemes
// @Description article_id content token  为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
// @Success 200 {string}  "{ "code": 200, "msg": "更新成功！" }"
// @Router  /user/addcommmits  [post]
func AddComment(c *gin.Context) {
	user_id := c.MustGet("Identification").(string)
	article_id := c.PostForm("article_id")
	parent_id := c.PostForm("parent_id")
	content := c.PostForm("content")
	if article_id == "" || content == "" {
		panic("必填参数不能为空！")
	}
	if f := models.GetCommentByArticleID(article_id, user_id); !f {
		panic("评论不存在！")
	}
	err := models.InsertComment(article_id, user_id, parent_id, content, utility.SetUuid())
	if err != nil {
		panic("评论失败，系统错误！" + err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "评论成功！",
	})

}
