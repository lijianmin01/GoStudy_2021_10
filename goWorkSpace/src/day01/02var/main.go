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
	fmt.Println(age)
}
