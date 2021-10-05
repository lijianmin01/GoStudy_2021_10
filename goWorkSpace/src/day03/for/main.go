package main

import "fmt"

// for 循环

// func main() {
// 	// 基本格式
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%d ", i)
// 	}

// 	fmt.Printf("\n")
// 	// 变种1
// 	var i = 0
// 	for ; i < 10; i++ {
// 		fmt.Printf("%d ", i)
// 	}
// 	i = 0
// 	// 变种2
// 	for i < 10 {
// 		fmt.Printf("%d ", i)
// 		i++
// 	}

// 	// for range 循环
// 	s := "Hello沙河"
// 	for i, v := range s {
// 		fmt.Printf("%d %c\n", i, v)
// 	}
// }

// 流程控制之跳出循环

func main() {
	// 当i=5的时候跳出for 循环
	// for i := 0; i < 10; i++ {
	// 	fmt.Print(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// 当i=5 的时候，跳过此次for循环（不执行打印）
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i)
	}
}
