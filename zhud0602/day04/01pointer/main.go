package main

import "fmt"

// func main() {
// 	a := 18
// 	b := &a
// 	fmt.Println(a)             //18
// 	fmt.Println(b)             //0xc000016098
// 	fmt.Printf("type:%T\n", b) //type:*int  Int类型的指针
// 	fmt.Println(&b)            //0xc000006028  存放指针地址的内存地址

// }

// func main() {
// 	//指针取值
// 	a := 10
// 	b := &a //取变量a的内存地址，将指针地址保存到b中
// 	c := &b
// 	d := &c
// 	fmt.Println(a) //10
// 	fmt.Println(b) //0xc000016098
// 	fmt.Println(c) //0xc000006028
// 	fmt.Println(d) //0xc000006030
// 	fmt.Println(*d)
// 	fmt.Println(*c)
// 	fmt.Println(*b)
// }

// // 10
// // 0xc000016098
// // 0xc000006028
// // 0xc000006030
// // 0xc000006028
// // 0xc000016098
// // 10

// func main() {
// 	var a *int  //在使用的时候不仅要声明，还要为他声明内存空间
// 	*a = 100
// 	fmt.Println(a)
// 	//invalid memory address or nil pointer dereference
// }

// func main() {
// 	a := new(int)
// 	b := new(string)
// 	c := new(bool)
// 	fmt.Printf("%T\n", a)
// 	fmt.Printf("%T\n", b)
// 	fmt.Printf("%T\n", c)

// 	fmt.Println(*a)
// 	fmt.Println(*b)
// 	fmt.Println(*c)

// }

// *int
// *string
// *bool
// 0
//
// false

func main() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(a)
	//0xc000016098
	fmt.Println(*a)
	//10
}
