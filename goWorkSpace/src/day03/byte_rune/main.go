package main

import "fmt"

func main() {
	//s := "Hello,ChinaTelecom~"
	// n := len(s)

	// // 循环
	// for i := 0; i < n; i++ {
	// 	fmt.Printf("%c ", s[i])
	// }
	// fmt.Printf("\n")
	// // 从字符串中拿出具体的字符
	// for _, c := range s {
	// 	fmt.Printf("%c", c)
	// }

	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转换成一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 把rune切片强制转换成字符串
}
