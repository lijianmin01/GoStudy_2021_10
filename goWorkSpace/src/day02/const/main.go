package main

import "fmt"

// 常量
/*
	定义常量后不能修改
	在程序运行期间不能修改
*/
const pi = 3.1415026

// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota: 类似枚举
// const 中每新增一行常量声明将使iota 计数一次
const (
	a1 = iota //0
	a2
	a3
	a4
	a5
)

// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3        //100
	c4        //100
)

const (
	d1 = iota //0
	d2 = 100  //100
	d3 = iota //2
	d4        //3
)

// 多个常量声明在一行
const (
	e1, e2 = iota + 1, iota + 2
	e3, e4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//pi=123
	// fmt.Println("n1:", n1)
	// fmt.Println("n2:", n2)
	// fmt.Println("n3:", n3)

	// fmt.Println("a1=", a1)
	// fmt.Println("a2=", a2)
	// fmt.Println("a3=", a3)
	// fmt.Println("a4=", a4)
	// fmt.Println("a5=", a5)

	// fmt.Println("c1", c1)
	// fmt.Println("c2", c2)
	// fmt.Println("c3", c3)
	// fmt.Println("c4", c4)

	// fmt.Println("e1", e1)
	// fmt.Println("e2", e2)
	// fmt.Println("e3", e3)
	// fmt.Println("e4", e4)

	fmt.Println(KB, MB, GB, TB, PB)
}
