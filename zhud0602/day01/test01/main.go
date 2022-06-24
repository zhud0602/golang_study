package main

import "fmt"

var (
	student_name string
	studentName  string
	StudentName  string

	//student_name string
	//不能出现同名的全局变量
)

func main() {

	//字符串声明1： var 变量名 string = ""
	//字符串声明2： var 变量名 = ""
	//字符串声明3： 变量名 := ""
	//字符串声明4： 变量名 = ""

	student_name = "abc"
	fmt.Println(student_name)
	//、全局变量可以和局部变量同名，但是在函数中会优先使用局部变量

	s2 := "20"
	fmt.Println(s2)
	s3 := "hahaha"
	fmt.Println(s3)
	var s4 string = "哈哈哈"
	fmt.Println(s4)

	var s5 = "aaa"
	fmt.Println(s5)

	s7 := "888"
	fmt.Println(s7)

	//s7 := "999"
	//fmt.Println(s7)
	//同一个作用域只能有同一个变量，否则会报错，无法使用

	//var s6:="666"
	//fmt.Println(s6)
	//报错  冒号和var不能同时用

	//变量离开作用域就不能使用
}
