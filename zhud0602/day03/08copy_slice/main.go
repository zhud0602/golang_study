package main

import "fmt"

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	c := make([]int, 6)
// 	fmt.Printf("c:%v\n", c)

// 	copy(c, a)
// 	fmt.Println(a)
// 	fmt.Println(c)

// 	c[0] = 100
// 	fmt.Println(a)
// 	fmt.Println(c)

// }

// [0 0 0 0 0 0]
// [1 2 3 4 5]
// [1 2 3 4 5 0]
// [1 2 3 4 5]
// [100 2 3 4 5 0]

// func main() {
// 	var a = make([]string, 5, 10)
// 	for i := 0; i < 10; i++ {
// 		a = append(a, "a")
// 	}
// 	fmt.Println(a)
// }

// func main() {
// 	var a = [...]int{3, 7, 8, 9, 1}
// 	sort.Interface(a)
// }

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1 := a1[:]
	s1 = append(s1, 19)
	//删掉索引为1的那个3
	s1[0] = 100
	//s1 = append(s1[0:1], s1[2:]...)
	fmt.Println(s1)
	//[1 5 7 9 11 13 15 17]
	fmt.Println(a1)
	//[1 5 7 9 11 13 15 17 17]

}
