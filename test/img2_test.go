package test

import (
	"github.com/fogleman/gg"
	"image/color"
	"log"
	"math/rand"
	"testing"
	"time"
)

func Test_1(t *testing.T) {
	width := 1920
	height := 1080
	lheight := 50
	s := "本周有4节课\n周1\n课程[计算机网络]\n节次[3--4]\n教室[智能518]\n周2\n课程[毛泽东思想和中国特色社会主义理论体系概论/习近平新时代中国特色社会主义思想概论]\n节次[1--2]\n教室[智能708]\n周3\n课程[Linux原理与应用]\n节次[1--2]\n教室[智能505]\n周3\n课程[Java Web开发技术]\n节次[7--8]\n教室[智能605]\n"
	rand.Seed(time.Now().UnixNano())

	//读取图片
	orange, err := gg.LoadImage("../img/o.png")
	if err != nil {
		log.Panicln(err)

	}
	//读取配置字体
	font, err := gg.LoadFontFace("../img/t.ttf", 60)
	if err != nil {
		log.Panicln(err)

	}
	f, err := gg.LoadFontFace("../img/5.ttf", 50)
	if err != nil {
		log.Panicln(err)
	}
	//创建容器
	d := gg.NewContext(width, height)
	//画矩形
	d.DrawRectangle(0, 0, float64(width), float64(height))

	//加载字体,颜色
	d.SetHexColor("#FF6000")
	d.Fill()
	d.SetFontFace(font)
	//d.SetColor(color.RGBA{255, 96, 0, 255})
	d.SetHexColor("#FAEDCD")
	d.DrawStringAnchored("课程表", float64(width/2), 50, 0.5, 0.5)
	d.DrawStringAnchored("---@lei", float64(width)-400, float64(height)-120, 0.5, 0.5)

	list := d.WordWrap(s, float64(width))
	for i, s2 := range list {
		d.DrawString(s2, 50, 200+float64(i)*float64(lheight))
	}
	d.SetColor(color.RGBA{245, 239, 231, 30})
	d.SetFontFace(f)

	//添加水印
	for i := 0; i < 10; i++ {
		d.Push()
		d.RotateAbout(50, float64(width)/2, float64(height)/2)
		d.DrawStringAnchored("@lei", float64(rand.Int63n(1920)), float64(rand.Int63n(1080)), 0.5, 0.5)
		d.Translate(50, 50)
		d.Pop()
	}

	d.DrawImageAnchored(orange, width-130, height-120, 0.5, 0.5)

	err = d.SavePNG("../img/lei.png")
	if err != nil {
		log.Panicln(err)
		return
	}

}
