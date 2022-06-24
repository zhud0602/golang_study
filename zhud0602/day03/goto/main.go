package main

import "fmt"

// func main() {
// 	var breakFlag bool
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if j == 2 {
// 				// 设置退出标签
// 				breakFlag = true
// 				break
// 			}
// 			fmt.Printf("%v-%v\n", i, j)
// 		}
// 		// 外层for循环判断
// 		if breakFlag {
// 			break
// 		}
// 	}
// }

//0-0
//0-1

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	//return
	// 标签
breakTag:
	//fmt.Println("结束for循环")
}
