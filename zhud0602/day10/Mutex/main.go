package main

import (
	"fmt"
	"sync"
)

var (
	// 计数器
	count int
	// count 变量的互斥锁
	countMux sync.Mutex
)

// 返回当前计数器的值
func Count() int {
	countMux.Lock()
	defer countMux.Unlock()
	return count
}

// 对计数器的值加一
func IncCount() {
	countMux.Lock()
	defer countMux.Unlock()
	count++
}
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	IncCount()
	count := Count()
	fmt.Println("Count =", count)
}
