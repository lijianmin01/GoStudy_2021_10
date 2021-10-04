package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOK bool

// 批量声明
var (
	name string
	age  int
	isOK bool
)

func main() {
	name = "lijianmin"
	age = 16
	isOK = true
	// Go语言所有的声明变量必须使用，没有使用的话，不会通过编译
	fmt.Println(isOK)
	fmt.Printf("name:%s", name)
	fmt.Println()
	fmt.Println(age)

	// 声明变量同时赋值
	var s1 string = "王思聪"
	fmt.Println(s1)

	// 类型推导
	var s2 = "20"
	fmt.Println(s2)

	// 简短变量声明
	s3 := "哈哈哈"
	fmt.Println(s3)
}
