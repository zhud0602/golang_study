package main

import "fmt"

// func main() {
// finger := 3
// switch finger {
// case 1:
// 	fmt.Println("大拇指")
// case 2:
// 	fmt.Println("食指")
// case 3:
// 	fmt.Println("中指")
// case 4:
// 	fmt.Println("无明指")
// case 5:
// 	fmt.Println("小指")
// default:
// 	fmt.Println("无效的输入!")
// }

// switch finger := 1; finger {
// case 1:
// 	fmt.Println("大拇指")
// case 2:
// 	fmt.Println("食指")
// case 3:
// 	fmt.Println("中指")
// case 4:
// 	fmt.Println("无明指")
// case 5:
// 	fmt.Println("小指")
// default:
// 	fmt.Println("无效的输入!")
// }   	//此finger 只能在switch函数中使用

// finger := 1
// switch {
// case finger == 1:
// 	fmt.Println("大拇指")
// case finger == 2:
// 	fmt.Println("食指")
// case finger == 3:
// 	fmt.Println("中指")
// case finger == 4:
// 	fmt.Println("无明指")
// case finger == 5:
// 	fmt.Println("小指")
// default:
// 	fmt.Println("无效的输入!")
// }

// switch score := 85; {
// case score >= 90:
// 	fmt.Println("优秀")
// case score >= 80:
// 	fmt.Println("良好")
// case score >= 60:
// 	fmt.Println("及格")
// case score >= 0:
// 	fmt.Println("不及格")
// default:
// 	fmt.Println("无效的输入!")
// }

// score := 20
// switch {
// case score >= 90:
// 	fmt.Println("优秀")
// case score >= 80:
// 	fmt.Println("良好")
// case score >= 60:
// 	fmt.Println("及格")
// case score >= 0:
// 	fmt.Println("不及格")
// default:
// 	fmt.Println("无效的输入!")
// }

// 	switch numbr := 70; numbr {
// 	case 1, 3, 5, 7, 9:
// 		fmt.Println("奇数")
// 	case 2, 4, 6, 8, 10:
// 		fmt.Println("偶数")
// 	default:
// 		fmt.Println(numbr)
// 	}
// }

// func main() {
// 	s := "a"
// 	switch {
// 	case s == "a":
// 		fmt.Println("a")
// 		fallthrough
// 	case s == "b":
// 		fmt.Println("b")
// 	case s == "c":
// 		fmt.Println("c")
// 	default:
// 		fmt.Println("...")
// 	}
// }

//a
//b

func main() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

//a
