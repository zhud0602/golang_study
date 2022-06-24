// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectType(x interface{}) {
// 	v := reflect.TypeOf(x)
// 	fmt.Printf("type:%v\n", v)
// }

// func main() {
// 	var a float32 = 3.14
// 	reflectType(a)
// 	var b int64 = 100
// 	reflectType(b)
// }

package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v  kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32
	var b myInt
	var c rune
	reflectType(a) //type:  kind:ptr
	reflectType(b) //type:myInt  kind:int64
	reflectType(c) //type:int32  kind:int32

	//Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空

	type person struct {
		name string
		age  int
	}

	type book struct{ title string }

	var d = person{
		name: "沙河小王子",
		age:  18,
	}

	var e = book{title: "《跟小王子学Go语言》"}

	reflectType(d) //type:person  kind:struct
	reflectType(e) //type:book  kind:struct

}
