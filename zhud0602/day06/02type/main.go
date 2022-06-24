package main

import "fmt"

//类型定义 自定义类型，编译完后保留
type NewInt int

//类型别名   只在代码编写中有效，编译完后无效
type MyInt = int

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	var c rune = '中'
	fmt.Printf("type of c:%T\n", c) //type of c:int32

	//int32的别名，几乎在所有方面等同于int32
	//它用来区分字符值和整数值
}
