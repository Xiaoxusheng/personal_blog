package main

import (
	"log"
	"personal_blog/router"
)

func main() {
	r := router.Router()

	err := r.Run(":8080")
	if err != nil {
		log.Panicln("star is err:" + err.Error())
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
