// // package main

// // import "fmt"

// // //Address 地址结构体
// // type Address struct {
// // 	Province string
// // 	City     string
// // }

// // //User 用户结构体
// // type User struct {
// // 	Name   string
// // 	Gender string
// // 	Address
// // }

// // // func main() {
// // // 	User1 := User{
// // // 		Name:   "小王子",
// // // 		Gender: "男",
// // // 		Address: Address{
// // // 			Province: "山东",
// // // 			City:     "威海",
// // // 		},
// // // 	}

// // // 	fmt.Printf("User1=%#v\n", User1)
// // // 	fmt.Printf("User1=%v\n", User1)
// // // }

// // func main() {
// // 	var User2 User
// // 	User2.Name = "小王子"
// // 	User2.Gender = "男"
// // 	User2.Address.Province = "山东"
// // 	User2.City = "威海"
// // 	fmt.Printf("User2=%#v\n", User2)
// // }

// // package main

// // type Address struct {
// // 	Province   string
// // 	City       string
// // 	CreateTime string
// // }

// // type Email struct {
// // 	Account    string
// // 	CreateTime string
// // }

// // type User struct {
// // 	Name   string
// // 	Gender string
// // 	Address
// // 	Email
// // }

// // func main() {
// // 	var User3 User
// // 	User3.Name = "沙河娜扎"
// // 	User3.Gender = "男"
// // 	User3.Address.CreateTime = "2019"
// // 	User3.Email.CreateTime = "2019"

// // }

// // package main

// // import "fmt"

// // //动物
// // type Animal struct {
// // 	Name string
// // }

// // func (a *Animal) move() {
// // 	fmt.Printf("%s会运动\n", a.Name)
// // }

// // type Dog struct {
// // 	feet uint8
// // 	*Animal
// // }

// // func (d *Dog) wang() {
// // 	fmt.Printf("%s会汪汪汪~~\n", d.Name)
// // }

// // func main() {
// // 	d1 := &Dog{
// // 		feet: 4,
// // 		Animal: &Animal{
// // 			Name: "乐乐",
// // 		},
// // 	}

// // 	fmt.Println(d1)
// // 	d1.wang()
// // 	d1.move()

// // }

// // package main

// // import "fmt"

// // func main() {
// // b := []byte("Hello Gopher!")
// // b[1] = 'T'
// // fmt.Println(b)
// // fmt.Println(string(b))

// // 	var data [10]byte

// // 	data[0] = 'A'

// // 	data[1] = 'B'

// // 	var str string = string(data[:])

// // 	fmt.Println(str)
// // }

package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i), //stu01 stu02
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Println(data)
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Students struct {
// 	RollNo []int
// 	Names  []string
// }

// func main() {
// 	//Create an object of Students structure.
// 	stu := &Students{
// 		RollNo: []int{101, 102, 103},
// 		Names:  []string{"Ravi", "Kavi", "Mavi"}}

// 	result, _ := json.Marshal(stu)
// 	fmt.Println(result)
// 	fmt.Println(string(result))
// }
