package main

import "fmt"

func main() {
	arr := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, value := range arr {
		sum = sum + value
	}
	fmt.Println(sum)
}

//求数组[1, 3, 5, 7, 8]所有元素的和  //24

// func main() {
// 	arr := [...]int{1, 3, 5, 7, 8}
// 	for i := 0; i < len(arr); i++ {
// 		for j := i + 1; j < 5; j++ {
// 			if arr[i]+arr[j] == 8 {
// 				fmt.Printf("(%d,%d)\n", i, j)
// 			}
// 		}
// 	}
// }

//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
