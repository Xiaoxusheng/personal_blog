package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"personal_blog/router"
)

func main() {
	f, _ := os.Create("gin.log")
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.ForceConsoleColor()
	r := router.Router()

	err := r.Run(":8080")
	if err != nil {
		log.Panicln("star is err:" + err.Error())
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
