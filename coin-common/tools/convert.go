package tools

import (
	"log"
	"strconv"
)

func ToFloat64(data string) float64 {
	// 将字符串转换为浮点数
	result, err := strconv.ParseFloat(data, 64) // 将字符串转换为 64 位浮点数
	if err != nil {
		log.Fatalf("Failed to convert to float64: %v", err) // 处理错误
	}
	return result
}

func ToInt64(data string) int64 {
	// 将字符串转换为 int64
	result, err := strconv.ParseInt(data, 10, 64) // 使用基数 10，转换为 64 位整数
	if err != nil {
		log.Fatalf("Failed to convert to int64: %v", err) // 处理错误
	}
	return result
}
