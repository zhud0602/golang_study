package main

import "fmt"

func main() {
	//十进制数
	a1 := 10
	fmt.Printf("%d\n", a1)
	fmt.Printf("%T\n", a1)

	// 八进制  以0开头 ,零+字母o
	var b = 0o377
	fmt.Printf("%o \n", b) // 377  ,八进制表示
	//fmt.Printf("%d \n", b) //7+7*8+3*8*8=7+56+192=255，十进制表示
	//fmt.Printf("%b \n", b) //11111111 ，二进制表示 377=[3][7][7]=[011][111][111]=011111111
	//fmt.Printf("%x \n", b) //ff ，十六进制表示

	//十六进制 ， 以0开头 ,零+字母x或者大写X
	var c = 0xff
	fmt.Printf("%d\n", c)
	//fmt.Printf("%b\n", c)
	//fmt.Printf("%0o\n", c)
	//fmt.Printf("%0x", c)

	e := 0xff
	fmt.Printf("%d\n", e) //255
	//fmt.Printf("%b\n", c)  //11111111
	//fmt.Printf("%0O\n", c) //0o377
	//fmt.Printf("%0o\n", c) //377
	//fmt.Printf("%0X\n", c) //FF
	//fmt.Printf("%0x\n", c) //ff

	d := int(0b00101101)
	fmt.Printf("%d\n", d)  //45
	fmt.Printf("%b\n", d)  //101101
	fmt.Printf("%0b\n", d) //101101
	fmt.Printf("%T\n", d)

	fmt.Printf("%0O\n", d) //0o55
	fmt.Printf("%O\n", d)  //0o55

	fmt.Printf("%o\n", d)  //55
	fmt.Printf("%0o\n", d) //55

	fmt.Printf("%0X\n", d) //2D
	fmt.Printf("%X\n", d)  //2D

	fmt.Printf("%0x\n", d) //2d
	fmt.Printf("%x\n", d)  //2d

	fmt.Printf("%T\n", d) //T必须大写

	f := int8(10)
	fmt.Printf("%d\n", f)
}
