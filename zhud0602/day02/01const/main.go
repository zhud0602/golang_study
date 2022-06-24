package main

import "fmt"

const pi = 3.1415926
const e = 2.7182

func main() {
	//fmt.Println(pi)

	//const (
	//	n1 = 100 
	//	n2
	//	n3
	//)

	//fmt.Println(n1)
	//fmt.Println(n2)
	//fmt.Println(n3)

	//	const (
	//		a1 = iota
	//		a2
	//		_
	//		a4
	//	)
	//	fmt.Println(a1)
	//	fmt.Println(a2)
	//	fmt.Println(a4)

	//const (
	//	b1 = iota //0
	//	b2 = 100  //100
	//	b3 = iota //2
	//	b4        //3
	//)
	//fmt.Println(b1)
	//fmt.Println(b2)
	//fmt.Println(b3)
	//fmt.Println(b4)

	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
