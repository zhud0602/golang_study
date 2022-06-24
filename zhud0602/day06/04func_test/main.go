// package main

// import "fmt"

// //dog 这是一个结构体
// type dog struct {
// 	name string
// }

// //方法
// func (d dog) wang() {
// 	fmt.Println("汪汪汪")
// }

// func main() {

// 	d1 := dog{
// 		name: "小黑",
// 	}

// 	d1.wang() //汪汪汪
// }

// package main

// import "fmt"

// type myint int

// func (m myint) hello() {
// 	fmt.Println("hello")
// }

// func main() {
// 	m := myint(100)
// 	m.hello()
// }

//hello

package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

// func main() {
// 	var (
// 		name    string
// 		age     int
// 		married bool
// 	)
// 	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
// 	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
// }

// func main() {
// 	var (
// 		name    string
// 		age     int
// 		married bool
// 	)
// 	fmt.Scanln(&name, &age, &married)
// 	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
// }
