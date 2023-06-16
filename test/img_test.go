package test

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func Test_img(T *testing.T) {
	// create a new context with the specified size
	width := 1080
	height := 1920
	dc := gg.NewContext(width, height)

	orange, err := gg.LoadImage("../img/o.png")
	// draw a rectangle
	dc.SetColor(color.RGBA{255, 96, 0, 255})
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()

	// draw text
	face, err := gg.LoadFontFace("../img/t.ttf", 60)
	if err != nil {
		log.Panicln(err)
		return
	}
	f, err := gg.LoadFontFace("../img/3.ttf", 200)

	str := "本周有4节课\n周1\n课程[计算机网络]\n节次[3--4]\n教室[智能518]\n周2\n课程[毛泽东思想和中国特色社会主义理论体系概论/习近平新时代中国特色社会主义思想概论]\n节次[1--2]\n教室[智能708]\n周3\n课程[Linux原理与应用]\n节次[1--2]\n教室[智能505]\n周3\n课程[Java Web开发技术]\n节次[7--8]\n教室[智能605]\n"
	dc.SetFontFace(face)
	dc.SetColor(color.RGBA{245, 239, 231, 255})
	//dc.DrawStringAnchored("leilong", float64(width)/2, float64(height)/2, 0.5, 0.5)
	//lens := dc.WordWrap(str, float64(width))
	//lheight := 80
	//for i, s := range lens {
	//	dc.DrawString(s+"\n", 50, 350+float64(i)*float64(lheight))
	//}
	dc.DrawStringWrapped(str, 00, 350, 0, 0, 100, 1, 0)
	l, h := dc.MeasureMultilineString(str, 1)
	fmt.Println("str", l, h)

	//dc.DrawStringWrapped(str, 50, 300, 0, 0, 1080, 1.2, 0)
	dc.SetFontFace(f)
	dc.DrawStringAnchored("Class Schedule", float64(width)/2, float64(80), 0.5, 0.5)
	dc.SetColor(color.RGBA{245, 239, 231, 30})
	//橘子图标
	dc.DrawImageAnchored(orange, width-120, height-140, 0.5, 0.5)

	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < 5; i++ {
		// 将绘图上下文状态保存到栈中，并进行旋转
		dc.Push()
		dc.RotateAbout(gg.Radians(40), 0, float64(height/2))
		dc.DrawStringAnchored("@Maving", float64(rand.Int63n(1000)), float64(rand.Int63n(1920)), 0.5, 0.5)
		dc.Pop()
		dc.Translate(50, 50)
	}
	// save the image to a file
	if err := dc.SavePNG("../img/1.jpg"); err != nil {
		fmt.Println(err)
	}

	dir, err := os.ReadDir("../img")
	if err != nil {
		log.Panicln(err)
		return
	}
	for _, k := range dir {
		fmt.Println(k.Name())

	}
	fmt.Println()

}
