package main

import "fmt"

type Mover interface {
	Move()
}

type Dog struct{}

func (d Dog) Move() {
	fmt.Println("狗会动")
}

type Cat struct{}

func (c *Cat) Move() {
	fmt.Println("猫会动")
}

func main() {
	var x Mover

	var d1 = Dog{}
	x = d1
	x.Move()

	var d2 = &Dog{}
	x = d2
	x.Move()
}
