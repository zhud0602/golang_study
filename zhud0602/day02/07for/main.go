package main

import "fmt"

func main() {
	//for循环
	//基本格式
	//for里面的变量i只能在当前循环中使用
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//for循环的初始语句可以被忽略，但是初始语句后的分号必须要写
	n := 0
	for ; n <= 10; n++ {
		fmt.Println(n)
	}

	//for循环的初始语句和结束语句都可以省略
	m := 0
	for m <= 5 {
		fmt.Println(m)
		m++
	}

	//无限循环
	// for {
	// 	a := 1
	// 	a++
	// 	if a > 5 {
	// 		break
	// 	}else{
	// 		fmt.Println(a)
	// 	}
	// }

	//for range 循环
	s := "hello沙河"
	fmt.Println(len(s)) //11  5+3+3 ,中文字符占3个字节

	for a, b := range s {
		fmt.Printf("%d %c\n", a, b)
	}

	//a是索引，b是对应索引的字符
	/*
	   0 h
	   1 e
	   2 l
	   3 l
	   4 o
	   5 沙
	   8 河
	*/

	for a, b := range s {
		fmt.Println(a, b)
	}
	/*
	   0 104
	   1 101
	   2 108
	   3 108
	   4 111
	   5 27801
	   8 27827
	*/

	for a := range s {
		fmt.Println(a)
	}
	/*
		0
		1
		2
		3
		4
		5
		8
	*/

	for _, b := range s {
		fmt.Printf("%c\n", b)
	}
	/*
		h
		e
		l
		l
		o
		沙
		河
	*/

}
