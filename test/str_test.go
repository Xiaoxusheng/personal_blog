package test

import (
	"fmt"
	"github.com/fogleman/gg"
	"testing"
)

func Test_str(t *testing.T) {
	dc := gg.NewContext(10, 10)
	h, _ := gg.LoadFontFace("../img/t.ttf", 50)
	dc.SetFontFace(h)
	str := "复的环境dkkfkf代码中使用了rune类型表示一个Unicode字符，函数返回一个布尔值表示该字符是否为中文。中文字符的Unicode编码范围为0x4E000x9FFF，因此我们可以通过比较字符的Uni"
	// 使用ContainsAny函数检查输入字符串是否包含英文字符
	var k float64 = 0
	f := ""
	list := make([]string, 0)
	for _, r := range str {
		width, _ := dc.MeasureString(string(r))
		f += string(r)
		k += width
		if k > 1080 {
			list = append(list, f)
			fmt.Println(f)
			k = 0
		}
		fmt.Printf("Character '%c' has width %f k %v \n", r, width, k)
	}

	//for len(str) > 0 {
	//	r, size := utf8.DecodeRuneInString(str)
	//	fmt.Println(string(r), size)
	//	f += string(r)
	//	k += size
	//	if k%45 == 0 {
	//		f += "\t"
	//	}
	//	str = str[size:]
	//	fmt.Println(f, k)
	//
	//}

}
