package main

import "fmt"

var i = initvar()

func initvar() int {
	fmt.Println("initvar...")
	return 100
}

func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println(i)
}

/*initvar...
init...
100*/
//执行顺寻：全局变量初始化>init()>main()
