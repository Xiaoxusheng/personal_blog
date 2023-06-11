package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_blog/models"
	"personal_blog/utility"
	"strconv"
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
//
// @Success 200 {string}   "{"code": 200,"data": {"list": [{"id": 1,"article_id": "52fd3b87-9572-4330-9371-9b7cf54628a2",	"user_id": "e5a6071b-baaf-45aa-a587-784d0ff9a575",	"parent_id": "","comment_id": "","content": "哈哈哈","status": 0,"created_time": "2023-06-09T17:26:30Z", "updated_time": "2023-06-09T17:26:30Z"	},{"id": 2,"article_id": "52fd3b87-9572-4330-9371-9b7cf54628a2", "user_id": "e5a6071b-baaf-45aa-a587-784d0ff9a575","parent_id": "","comment_id": "e725fe87-a881-4cc9-ad7a-bc48ccdb97dd","content": "哈哈哈", "status": 0,"created_time": "2023-06-09T17:46:38Z", "updated_time": "2023-06-09T17:46:38Z"	}, ] }, "msg": "获取数据成功！"} "
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
// @Success 200 {string}  "{"code":200,"msg":"评论成功！"}"
// @Router  /user/addcomments  [post]
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

// PingExample godoc
// @Summary 审核评论接口
// @title My awesome API
// @version 1.0
// @host localhost:8080
// @Param comment_id query string true "评论唯一标识"
// @Param status query string true "审核状态 0 表示未审核 1 表示审核通过 2 表示审核不通过"
// @Param Authorization header string true "token"
// @Schemes
// @Description comment_id status content token  为必填
// @Tags 私有方法
// @Accept json
// @Produce json
// @Success 200 {string}    "{ "code": 200, "msg": "审核成功！" }"
// @Router  /api/examinecomment  [get]
func ExamineComment(c *gin.Context) {
	//评论唯一id
	comment_id := c.Query("comment_id")
	status := c.Query("status")
	if comment_id == "" || !utility.Contains(utility.StatusList, status) {
		panic("参数不正确！")
	}
	atoi, err := strconv.Atoi(status)
	if err != nil {
		return
	}
	if err = models.GetByCommentId(comment_id); err != nil {
		panic("评论不存在！")
	}
	err = models.UpdateComment(comment_id, atoi)
	if err != nil {
		panic("评论审核失败！")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "审核成功！",
	})
}
