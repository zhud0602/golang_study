package main

import (
	"fmt"
)

func main() {
	f1 := 1.234567
	fmt.Printf("%T\n", f1) //默认go语言中的小数都是float64类型

	const f3 = 1.234555
	fmt.Printf("%T\n", f3)

	f2 := float32(1.234567)
	fmt.Printf("%T\n", f2)

	//f2 = f1 //不同类型的浮点数不能直接赋值

}
