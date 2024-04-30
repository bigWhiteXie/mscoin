package validate

import (
	"regexp"
)

func IsPhoneFomatter(phoneNumber string) bool {
	// 定义手机号的正则表达式
	// 定义手机号的精确正则表达式
	pattern := `^1[3-9]\d{9}$`
	// 编译正则表达式
	regex := regexp.MustCompile(pattern)

	// 使用正则表达式进行匹配
	return regex.MatchString(phoneNumber)
}
