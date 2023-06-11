package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"personal_blog/db"
	"strconv"
)

// 上传文件
// PingExample godoc
// @Summary  上传文件接口
// @Schemes
// @Param file formData file true "表单name"
// @Param token header string true "token"
// @Description 上传文件
// @Description file token为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{"code":1,"msg":"\u0001个文件上传成功","url":"127.0.0.1:8080/img/12.png"}
// @Router  /user/file      [post]
func File(c *gin.Context) {
	var filename string
	key := "img"
	ctx := context.Background()
	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	go func() {
		_ = os.Mkdir("img", 0750)
	}()
	if err != nil {
		return
	}
	files := form.File["file"]
	_, err = db.Rdb.Exists(ctx, key).Result()
	if err != nil {
		return
	}

	for _, file := range files {
		//查看file.Filename是否存在
		result, err := db.Rdb.SIsMember(ctx, key, file.Filename).Result()
		if err != nil {
			continue
		}
		if result {
			panic("文件已经存在，请不要重复上传！")
		}
		db.Rdb.SAdd(ctx, key, file.Filename)
		//上传指定目录
		err = c.SaveUploadedFile(file, "./img/"+file.Filename)
		if err != nil {
			panic(err)
		}
		filename = file.Filename
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  strconv.Itoa(len(files)) + "个文件上传成功",
		"code": 200,
		"data": gin.H{
			"url": "http://116.198.44.154:8080/img/" + filename,
		},
	})
}
