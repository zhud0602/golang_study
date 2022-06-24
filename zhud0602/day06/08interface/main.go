// package main

// import "fmt"

// type dog struct{}

// type cat struct{}

// type person struct{}

// func (d dog) speak() {
// 	fmt.Println("汪汪汪~!")
// }

// func (c cat) speak() {
// 	fmt.Println("喵喵喵")
// }

// func (p person) speak() {
// 	fmt.Println("啊啊啊~!")
// }

// func (p person) up() {
// 	fmt.Println("eee~!")
// }

// type speaker interface {
// 	speak()
// }

// func da(x speaker) {
// 	x.speak()

// }

// func main() {
// 	var d1 dog
// 	var c1 cat
// 	var p1 person
// 	da(d1)
// 	da(c1)
// 	da(p1)

// }

// //接口是一种类型，凡是含有接口中方法的类型，都可以作为此接口变量

package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动!")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s\n", food)
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {
	var a1 animal
	bc := cat{
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)

	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	fmt.Printf("%T\n", a1)
}
