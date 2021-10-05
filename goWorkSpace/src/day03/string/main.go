package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "\"D:\\Anaconda3\\condabin\""
	fmt.Println(path)

	// 多行的字符串
	s2 := `
		世情薄
			人情恶
	雨落黄昏花遗落
	`
	fmt.Println(s2)

	s3 := "D:\\Anaconda3\\condabin"
	s4 := "\\test.py"
	// 字符串相关操作
	fmt.Println(len(s3))
	fmt.Println(s3 + s4)

	ss1 := fmt.Sprintf("%s%s", s3, s4)
	fmt.Println(ss1)

	// 分隔
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(s3, "E:"))
	fmt.Println(strings.Contains(s3, "D:"))

}
