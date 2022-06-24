package main

import "fmt"

func main() {
	age := 19

	//if条件判断
	if age > 18 {
		fmt.Println("学习golang语言中")
	} else {
		fmt.Println("否则执行这个代码")
	}

	//多个判断条件
	score := 65

	if score >= 60 {
		fmt.Println("及格")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 90 {
		fmt.Println("优秀")
	} else {
		fmt.Println("不及格")
	}

	//可在IF判断表达式之前添加执行语句
	//age变量此时只在if条件判断语句中生效
	if score2 := 52; score2 >= 90 {
		fmt.Println("优秀")
	} else if score2 >= 80 {
		fmt.Println("良好")
	} else if score2 >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// fmt.Println(score2) //在这里是找不到age

	





}
