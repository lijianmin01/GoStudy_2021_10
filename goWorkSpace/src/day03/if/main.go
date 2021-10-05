package main

import "fmt"

func main() {
	// if 判断
	age := 21
	if age > 18 {
		fmt.Println("雄安我来了")
	} else {
		fmt.Println("继续加油")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("天天向上")
	}
}
