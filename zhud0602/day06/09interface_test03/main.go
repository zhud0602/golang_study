package main

import "fmt"

type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

type Dog struct {
	name string
}

func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪~!\n", d.name)
}

func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	// var d Dog
	// d.name = "旺财"
	d := &Dog{"旺财"}
	var s Sayer = d
	var m Mover = d
	s.Say()
	m.Move()

}

