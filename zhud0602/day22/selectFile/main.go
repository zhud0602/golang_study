package main

import "fmt"

type cat struct {
}

type dog struct {
}

type pet interface {
	eat()
	sleep()
}

func (d dog) eat() {
	fmt.Println("dog eat ...")
}

func (d dog) sleep() {
	fmt.Println("dog sleep ...")
}

func (c cat) eat() {
	fmt.Println("cat eat ...")
}

func (c cat) sleep() {
	fmt.Println("cat sleep ...")
}

type person struct {
	name string
}

func (p1 person) care(p2 pet) {
	p2.sleep()
	p2.eat()
}

func main() {
	dog := dog{}
	cat := cat{}
	person := person{}
	person.care(dog)
	person.care(cat)
}
