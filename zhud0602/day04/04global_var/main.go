package main

import "fmt"

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

// func sub(x, y int) int {
// 	return x - y
// }

func main() {
	var c calculation = add         // 声明一个calculation类型的变量c  把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 3  像调用add一样调用c

	f := add                        // 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 30  像调用add一样调用f
}
