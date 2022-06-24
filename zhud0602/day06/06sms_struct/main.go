package main

import (
	"fmt"
	"os"
)

var smr sms

func showMenu() {
	fmt.Println("----------welcome sms!-----------")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
}

func main() {
	smr = sms{
		allStudent: make(map[int64]student, 100),
	}

	//对sms赋值初始化
	//对map初始化

	for {
		showMenu()
		//等待用户输入选项
		fmt.Print("请输入序号")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：", choice)

		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.alterStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("滚")
		}
	}
}
