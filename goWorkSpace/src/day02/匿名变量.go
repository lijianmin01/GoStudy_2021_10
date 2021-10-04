package 匿名变量

import "fmt"

func foo(int, string) {
	return 10, "123A"
}

func main() {
	x, _ = foo()
	fmt.Println("x=", x)
	_, y = foo()
	fmt.Println("y=", y)
}
