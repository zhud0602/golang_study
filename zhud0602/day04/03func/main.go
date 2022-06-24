package main

import "fmt"

//insum函数名称 ，x,y参数
// func insum(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(insum(1, 2)) //3
// }

// func sayhello() {
// 	fmt.Println("hello 沙河")
// }

// func main() {
// 	sayhello()
// }

//hello 沙河  调用有返回值的函数时，可以不接收其返回值。

//函数的参数中如果相邻变量的类型相同，则可以省略类型，例如:
// func insum(x, y int, m, n string) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(insum(1, 2, "a", "b")) //3

// }

//可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
// func intSum2(x ...int) int {
// 	fmt.Println(x) //x是一个切片
// 	sum := 0
// 	for _, value := range x {
// 		sum = sum + value
// 	}
// 	return sum
// }

// func main() {
// 	fmt.Println(intSum2(1, 3, 5, 7, 9))
// }

// [1 3 5 7 9]
// 25

// 固定参数搭配可变参数使用时，可变参数要放在固定参数的后面，示例代码如下：
// func intSum3(x int, y ...int) int {
// 	fmt.Println(x, y)
// 	sum := x
// 	for _, v := range y {
// 		sum = sum + v
// 	}
// 	return sum
// }

// func main() {
// 	ret5 := intSum3(100)                //100 []
// 	ret6 := intSum3(100, 10)            //100 [10]
// 	ret7 := intSum3(100, 10, 20)        //100 [10 20]
// 	ret8 := intSum3(100, 10, 20, 30)    //100 [10 20 30]
// 	fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160

// }

//函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
// func calc(x, y int) (sum, sub int) {
// 	sum = x + y
// 	sub = x - y
// 	return sum, sub
// }

// func main() {
// 	fmt.Println(calc(20, 10)) //30 10
// }

//当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
func somefunc(x string) []int {
	if x == "" {
		return nil
	} else {
		return []int{1, 2}
	}
}

func main() {
	fmt.Println(somefunc(""))   //[]
	fmt.Println(somefunc("AA")) //[1 2]
}
