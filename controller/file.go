package controller

import (
	"context"
	"github.com/fogleman/gg"
	"github.com/gin-gonic/gin"
	"image/color"
	"log"
	"math/rand"
	"net/http"
	"os"
	"personal_blog/db"
	"strconv"
	"time"
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

func CreateImg(c *gin.Context) {
	imgname := c.Query("imgname")

	// create a new context with the specified size
	width := 1080
	height := 1920
	dc := gg.NewContext(width, height)

	orange, err := gg.LoadImage("./img/o.png")
	// draw a rectangle
	dc.SetColor(color.RGBA{255, 96, 0, 255})
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()

	// draw text
	face, err := gg.LoadFontFace("./img/t.ttf", 60)
	if err != nil {
		log.Panicln(err)
		return
	}
	f, err := gg.LoadFontFace("./img/3.ttf", 300)

	dc.SetFontFace(face)
	dc.SetColor(color.RGBA{245, 239, 231, 255})
	//dc.DrawStringAnchored("leilong", float64(width)/2, float64(height)/2, 0.5, 0.5)
	lens := dc.WordWrap(imgname, float64(width))
	lheight := 80
	for i, s := range lens {
		dc.DrawString(s+"\n", 0, 300+float64(i)*float64(lheight))
	}

	dc.SetFontFace(f)
	dc.DrawStringAnchored("IPhone", float64(width)/2, float64(80), 0.5, 0.5)
	dc.SetColor(color.RGBA{245, 239, 231, 30})
	//橘子图标
	dc.DrawImageAnchored(orange, width-120, height-140, 0.5, 0.5)

	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < 5; i++ {
		dc.Push()
		dc.RotateAbout(gg.Radians(40), 0, float64(height/2))
		dc.DrawStringAnchored("@Maving", float64(rand.Int63n(1000)), float64(rand.Int63n(1920)), 0.5, 0.5)
		dc.Pop()
		dc.Translate(50, 50)
	}
	// save the image to a file
	if err := dc.SavePNG("./img/1.jpg"); err != nil {
		panic(err)
	}
}
