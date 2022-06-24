package main

import "fmt"

// func f1() int {
// 	x := 5
// 	defer func() {
// 		x++
// 	}()
// 	return x //1.返回值赋值x=5, 执行defer:x=x+1=6
// }

// func f2() (x int) {
// 	defer func() {
// 		x++
// 	}()
// 	return 5 //1.先给返回值赋值x=5,执行defer:x=x+1=6 ,返回x:6
// }

// func f3() (y int) {
// 	x := 5
// 	defer func() {
// 		x++
// 	}()
// 	return x
// }
// func f4() (x int) {
// 	defer func(x int) {
// 		x++
// 	}(x)
// 	return 5
// }
// func main() {
// 	fmt.Println(f1()) //5
// 	fmt.Println(f2()) //6
// 	fmt.Println(f3()) //5
// 	fmt.Println(f4()) //5
// }

// func calc(index string, a, b int) int {
// 	ret := a + b
// 	fmt.Println(index, a, b, ret)
// 	return ret
// }

// func main() {
// 	x := 1
// 	y := 2
// 	defer calc("AA", x, calc("A", x, y))
// 	x = 10
// 	defer calc("BB", x, calc("B", x, y))
// 	y = 20
// }

//执行顺序：
//1.defer中有嵌套函数，先执行最里面，calc("A",1,2) ,返回A,1,2,3
//2.执行defer calc("AA", 1, 3)
//3.执行calc("B",10,2) ,返回B,10,2,12
//4.执行defer calc("BB",10,12)
//5.执行calc("BB",10,12) ,返回BB,10,12,22
//6.执行calc("AA",1,3) ,返回AA,1,3,4

func main() {
	fmt.Printf("%2.10s", "lixiang")
}
