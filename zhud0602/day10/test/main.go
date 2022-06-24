// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name   string   `json:"name" describe:"姓名"`
// 	Age    int      `json:"age" describe:"年龄"`
// 	Gender bool     `json:"gender" describe:"性别"`
// 	Hobby  []string `json:"hobby" describe:"爱好"`
// }

// func main() {
// 	//实例化结构体
// 	var s1 = Student{
// 		Name:   "张三",
// 		Age:    18,
// 		Gender: true,
// 		Hobby:  []string{"吃", "喝", "pia", "玩"},
// 	}
// 	var t = reflect.TypeOf(s1)
// 	fmt.Println(t.Name())     //Student
// 	fmt.Println(t.Kind())     //struct
// 	fmt.Println(t.NumField()) //结果:4,表示多少个字段
// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i) //每个结构体对象
// 		/*
// 			{Name  string json:"name" describe:"姓名" 0 [0] false}
// 			{Age  int json:"age" describe:"年龄" 16 [1] false}
// 			{Gender  bool json:"gender" describe:"性别" 24 [2] false}
// 			{Hobby  []string json:"hobby" describe:"爱好" 32 [3] false}
// 		*/
// 		//fmt.Println(field)
// 		fmt.Println("------")
// 		fmt.Printf("field.Name:%v\n", field.Name)
// 		fmt.Printf("field.Index:%v\n", field.Index)
// 		fmt.Printf("field.Type:%v\n", field.Type)
// 		fmt.Printf("field.Tag:%v\n", field.Tag.Get("describe"))

// 	}

// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Student struct {
// 	Name   string   `json:"name" describe:"姓名"`
// 	Age    int64    `json:"age" describe:"年龄"`
// 	Gender bool     `json:"gender" describe:"性别"`
// 	Hobby  []string `json:"hobby" describe:"爱好"`
// }

// func main() {
// 	//实例化结构体
// 	var s1 = Student{
// 		Name:   "张三",
// 		Age:    18,
// 		Gender: true,
// 		Hobby:  []string{"吃", "喝", "pia", "玩"},
// 	}
// 	fmt.Println(s1)
// 	var t = reflect.TypeOf(s1)
// 	fmt.Println(t.Name())     //Student
// 	fmt.Println(t.Kind())     //struct
// 	fmt.Println(t.NumField()) //结果:4,表示多少个字段
// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i) //每个结构体对象
// 		/*
// 			{Name  string json:"name" describe:"姓名" 0 [0] false}
// 			{Age  int json:"age" describe:"年龄" 16 [1] false}
// 			{Gender  bool json:"gender" describe:"性别" 24 [2] false}
// 			{Hobby  []string json:"hobby" describe:"爱好" 32 [3] false}
// 		*/
// 		fmt.Println(field)
// 		fmt.Println(reflect.TypeOf(field).Kind())
// 		fmt.Println("------")
// 		fmt.Printf("field.Name:%v\n", field.Name)
// 		fmt.Printf("field.Index:%v\n", field.Index)
// 		fmt.Printf("field.Type:%v\n", field.Type)
// 		fmt.Printf("field.Tag:%v\n", field.Tag.Get("describe"))
// 		fmt.Printf("field.Interface:%v\n", field.Interface())

// 	}
// }

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string   `json:"name" describe:"姓名"`
	Age    int64    `json:"age" describe:"年龄"`
	Gender bool     `json:"gender" describe:"性别"`
	Hobby  []string `json:"hobby" describe:"爱好"`
}

func main() {
	//实例化结构体
	var s1 = Student{
		Name:   "张三",
		Age:    18,
		Gender: true,
		Hobby:  []string{"吃", "喝", "pia", "玩"},
	}

	var m = reflect.ValueOf(s1)
	for i := 0; i < m.NumField(); i++ {
		field2 := m.Field(i)
		fmt.Println("------")
		fmt.Printf("Kind:%v\n", field2.Kind())
		fmt.Printf("值:%v\n", field2.Interface())
	}
}
