package main

import (
	"fmt"
	"time"
)

func printInfo(s string) {
	for {
		fmt.Println("Info =", s)
		// 延时1秒
		time.Sleep(time.Second)
	}
}
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	// 使用普通函数创建 goroutine
	go printInfo("HaiCoder Golang")
	time.Sleep(time.Duration(3000) * time.Millisecond)
}
