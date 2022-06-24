package main

import (
	"fmt"
)

// func f2(x, y int) {
// 	fmt.Println("this is f2")
// 	fmt.Println(x + y)
// }

// //定义一个函数对f2进行封装
// func lixiang(x func(int, int), m, n int) func() {
// 	temp := func() {
// 		x(m, n)
// 	}
// 	return temp
// }

// func main() {
// 	ret := lixiang(f2, 100, 200) //返回一个函数： ret := func(){f2(100, 200)}
// 	ret()                        //实际执行函数f2(100,200) ，this is  f2   300
// }

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y //x=x+y
		return x
	}
}
func main() {
	var f = adder()
	fmt.Println(f(10)) //10   x=0, y=10
	fmt.Println(f(20)) //30   x=0, y=10
	fmt.Println(f(30)) //60   x=0, y=10
}

// 	f1 := adder()
// 	fmt.Println(f1(40)) //40   x=0, y=10
// 	fmt.Println(f1(50)) //90   x=0, y=10
// }

// func adder2(x int) func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }
// func main() {
// 	var f = adder2(10) //x=10 ,返回值f=func(y int)int{}
// 	fmt.Println(f(10)) //20     传参y=10,返回值x=x+y=10+10=20
// 	fmt.Println(f(20)) //40		此时x=20,传参y=20,返回值x=x+y=20+20=40
// 	fmt.Println(f(30)) //70     此时x=40,传参y=30,返回值x=x+y=20+20=70

// 	f1 := adder2(20)    //x=20, 返回值f=func(y int)int{}
// 	fmt.Println(f1(40)) //60    传参y=40,返回值x=x+y=20+40=60
// 	fmt.Println(f1(50)) //110   此时x=60,传参y=50,返回值x=x+y=60+50=110
// }

// func makeSuffixFunc(suffix string) func(string) string {
// 	return func(name string) string {
// 		if !strings.HasSuffix(name, suffix) {
// 			return name + suffix
// 		}
// 		return name
// 	}
// }

// func main() {
// 	jpgFunc := makeSuffixFunc(".jpg") //传参suffix=".jpg"，返回值jpgFunc :=func(name string) string
// 	txtFunc := makeSuffixFunc(".txt") //传参suffix=".txt"，返回值jpgFunc :=func(name string) string
// 	fmt.Println(jpgFunc("test"))      //test.jpg    //传参name="test",变量suffix=".jpg"  返回值name + suffix="test.jpg"
// 	fmt.Println(txtFunc("test"))      //test.txt    //传参name="test",变量suffix=".txt"  返回值name + suffix="test.txt"
// }

// func calc(base int) (func(int) int, func(int) int) {
// 	add := func(i int) int {
// 		base += i
// 		return base
// 	}

// 	sub := func(i int) int {
// 		base -= i
// 		return base
// 	}
// 	return add, sub
// }

// func main() {
// 	f1, f2 := calc(10)        // 传参base=10 ,返回值f1=add, f2=sub ;
// 	fmt.Println(f1(1), f2(2)) //11 9  f1(1) ,i=1,base=10 ;返回值base=11   f(2), i=2 ,base=11 ;返回值base=11-2=9
// 	fmt.Println(f1(3), f2(4)) //12 8  f1(3) ,i=3,base=9 ;返回值base=12   f(4), i=4 ,base=12 ;返回值base=11-2=8
// 	fmt.Println(f1(5), f2(6)) //13 7  f1(5) ,i=5,base=8 ;返回值base=13   f(6), i=6 ,base=13 ;返回值base=13-6=7
// }
