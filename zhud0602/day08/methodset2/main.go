package main

import (
	"fmt"
)

//自定义List类型:
type List []int

//Len方法（接收者List的值类型/指针类型）
func (l List) Len() int {
	return len(l)
}

//Append方法（接收者List的指针类型）
func (l *List) Append(val int) {
	*l = append(*l, val)
}

//定义Appender接口，它实现了Append(int)方法
type Appender interface {
	Append(int)
}

//定义了函数CountInto
func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

//定义了接口Lener，实现Len()int方法
type Lener interface {
	Len() int
}

//定义了函数
func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	// A bare value
	var lst List
	fmt.Println(lst)
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	// CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID: Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) // VALID: Identical receiver type
	fmt.Println(plst)
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
