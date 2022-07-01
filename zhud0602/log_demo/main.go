package main

import (
	"fmt"
	"log"
)

func test1() {
	defer fmt.Println("结束前执行")
	log.Print("my log")
	log.Panic("my panic")
	fmt.Println("end ...")
}

func test2() {
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 18
	log.Println(name, "+", age)
}

func test3() {
	defer fmt.Println("defer...")
	log.Print("my log")
	log.Fatal("fatal...")
	//os.exit(1)  log.Fatal 并不会执行defer函数
	fmt.Println("end...")
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}

func main() {
	i := log.Flags()
	fmt.Printf("i:%v/n", i)
	log.Print("my log...")

}
