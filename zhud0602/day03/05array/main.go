package main

import "fmt"

// func main() {
// 	b1 := [3]int{1, 2, 3}
// 	b2 := b1
// 	b2[0] = 100
// 	fmt.Println(b1, b2)
// }

//[1 2 3] [100 2 3]

// func main() {
// 	b1 := []int{1, 2, 3}
// 	b2 := b1
// 	b2[0] = 100
// 	fmt.Println(b1, b2)
// }

//[100 2 3] [100 2 3]

// func main() {
// 	b1 := []int{1, 2, 3}
// 	b2 := b1
// 	fmt.Println(b1, b2)
// }

//[1 2 3] [1 2 3]

// func main() {
// 	b1 := []int{1, 2, 3}
// 	b1[2] = 4
// 	fmt.Println(b1)
// }
// [1 2 4]

//根据索引来初始化数组
// func main() {
// 	b1 := [3]int{0: 1, 2: 3}
// 	fmt.Println(b1)
// }

//[1 0 3]，默认是0

//根据索引来初始化数组
// func main() {
// 	b1 := []int{0: 1, 9: 3}
// 	fmt.Println(b1)
// }

//[1 0 0 0 0 0 0 0 0 3]

//数组的遍历
// func main() {
// 	city := []string{"北京", "上海", "深圳"}
// 	for i := 0; i < len(city); i++ {
// 		fmt.Println(city[i])
// 	}
// }

/*
北京
上海
深圳
*/

// func main() {
// 	city := []string{"北京", "上海", "深圳"}
// 	for i, j := range city {
// 		fmt.Printf("[%d]%s\n", i, j)
// 	}
// }
/*
[0]北京
[1]上海
[2]深圳
*/

// func main() {
// 	citys := [...]string{"北京", "上海", "深圳"}
// 	city0 := citys[0]
// 	fmt.Println(city0)
// 	for _, j := range citys {
// 		fmt.Println(j)
// 	}
// }

// 北京
// 北京
// 上海
// 深圳

//二维数组
// func main() {
// 	citys := [3][2]string{
// 		{"北京", "上海"},
// 		{"广州", "深圳"},
// 		{"武汉", "成都"},
// 	}

// 	for index, value := range citys {
// 		fmt.Println(index, value)
// 	}

// 	fmt.Println(citys[2])
// 	fmt.Println(citys[2][0])
// 	fmt.Println(citys)
// }

// 0 [北京 上海]
// 1 [广州 深圳]
// 2 [武汉 成都]
// [武汉 成都]
// 武汉
//[[北京 上海] [广州 深圳] [武汉 成都]]

// func main() {
// 	a1 := [...][2]int{
// 		{1, 2},
// 		{3, 4},
// 		{5, 6},
// 	}
// 	fmt.Println(a1)
// }

//Go语言中的函数传递的都是值（Ctrl+c,Ctrl+v）
// func f1(a [3]int) [3]int {
// 	a[1] = 100 //此处修改的是副本
// 	return a
// }

// func main() {
// 	x := [3]int{1, 2, 3}
// 	fmt.Println(f1(x)) //[1 100 3]   返回的是副本的数组
// 	fmt.Println(x)     //[1 2 3]      返回的是原来数组
// }

//切片不存值，s1 s2指向的是同一个底层数组
// func main() {
// 	s1 := []int{1, 2, 3} //[1 2 3]
// 	s2 := s1
// 	fmt.Println(s2)
// 	s2[1] = 200
// 	fmt.Println(s2) //[1 200 3]
// 	fmt.Println(s1) //[1 200 3]
// }
// func main() {
// 	var a2 []int //应该将变量声明与下一行的赋值合并
// 	a2 = make([]int, 1)
// 	a2[0] = 100
// 	fmt.Println(a2)
// }

//[100]

func main() {
	a1 := []int{1, 3, 5, 7, 9}
	a2 := a1
	a3 := make([]int, 5)
	copy(a3, a1) ////使用copy()函数将切片a1中的元素复制到切片a3
	a2[0] = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	// [100 3 5 7 9]
	// [100 3 5 7 9]
	// [1 3 5 7 9]

}
