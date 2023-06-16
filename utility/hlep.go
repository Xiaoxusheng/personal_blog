package utility

import (
	"strings"
)

//type ResponseError struct {
//	Error string `json:"error"`
//	Code  int
//}

var MySigningKey = []byte("my_bl^%%^og_84775M")

var List = []string{"0", "1"}
var StatusList = []string{"0", "1", "2"}

func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func AddTabsAtInterval(input string, interval int) string {
	var parts []string
	englishAlphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 使用ContainsAny函数检查输入字符串是否包含英文字符
	strings.ContainsAny(input, englishAlphabet)

	// 将输入字符串按照指定的间隔进行分割

	for i := 0; i < len(input); i += interval {
		end := i + interval

		if end > len(input) {
			end = len(input)
		}
		parts = append(parts, input[i:end])
		//fmt.Println(end, parts)
	}

	// 在每个分段末尾加上制表符
	for i := 1; i < len(parts); i++ {
		parts[i] = "\t" + parts[i]
	}

	// 使用Join函数将各个分段拼接成最终的字符串
	return strings.Join(parts, "")
}
