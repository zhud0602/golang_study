package main

import "fmt"

func main() {
	b1 := true
	var b2 bool //默认是false
	var b3 bool = false
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value:%v\n", b2, b2)
	fmt.Printf("%T\n", b3)

}
