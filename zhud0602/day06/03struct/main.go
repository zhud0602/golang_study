// package main

// import "fmt"

// //结构体
// type person struct {
// 	name   string
// 	age    int
// 	gender string
// 	hobby  []string
// }

// func main() {
// 	//声明一个person类型的变量zhoulin
// 	var zhoulin person
// 	//通过字段赋值
// 	zhoulin.name = "周林"
// 	zhoulin.age = 18
// 	zhoulin.gender = "女"
// 	zhoulin.hobby = []string{"游泳", "篮球", "足球"}
// 	fmt.Println(zhoulin)        //{周林 18 女 [游泳 篮球 足球]}
// 	fmt.Printf("%T\n", zhoulin) //main.person
// 	fmt.Println(zhoulin.name)   //周林

// 	var p person
// 	p.name = "黄州"
// 	p.age = 20
// 	fmt.Printf("type:%T value:%v\n", p, p) //type:main.person value:{黄州 20  []}
// }

// package main

// import "fmt"

// type person struct {
// 	name   string
// 	gender string
// }

// //go语言中函数传参数永远传的是拷贝
// func f(x person) {
// 	x.gender = "男" //修改的是传参的副本
// }

// func f2(x *person) {
// 	//(*x).gender = "男" //根据内存地址找到那个原来变量，修改的是原来变量 ，语法糖：(*x).gender = x.gender
// 	x.gender = "男"
// }

// func main() {
// 	var p person
// 	p.name = "周林"
// 	p.gender = "女"
// 	f(p)
// 	fmt.Println(p.gender) // 女
// 	f2(&p)
// 	fmt.Println(p.gender) //男
// }
// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int8
// }

// func main() {
// 	var p1 person
// 	p1.name = "沙河娜扎"
// 	p1.city = "北京"
// 	p1.age = 18
// 	fmt.Printf("p1=%v\n", p1)          //p1={沙河娜扎 北京 18}
// 	fmt.Printf("p1=%#v\n", p1)         //p1=main.person{name:"沙河娜扎", city:"北京", age:18}
// 	fmt.Printf("p1.name=%#v", p1.name) //p1.name="沙河娜扎"
// }

// package main

// import "fmt"

// func main() {
// 	type person struct {
// 		name string
// 		city string
// 		age  int
// 	}

// 	var p2 = new(person)
// 	fmt.Printf("%p\n", &p2)  //0xc000006028
// 	fmt.Printf("%T\n", p2)   //*main.person
// 	fmt.Printf("p2=%#v", p2) //p2=&main.person{name:"", city:"", age:0}  new func(type)*type
// }

// package main

// import "fmt"

// func main() {
// 	type person struct {
// 		name string
// 		city string
// 		age  int
// 	}

// }
// package main

// import "fmt"

// type student struct {
// 	name string
// 	age  int
// }

// func main() {
// 	m := make(map[string]*student)
// 	stus := []student{
// 		{name: "小王子", age: 18},
// 		{name: "娜扎", age: 23},
// 		{name: "大王八", age: 9000},
// 	}

// 	for _, stu := range stus {
// 		m[stu.name] = &stu //一个变量只能存在一个内存地址中，变量值可以变，但是地址是一直不变的
// 		fmt.Printf("%p\n", &stu)
// 		fmt.Println(stu)
// 	}

// 0xc000004078
// {小王子 18}
// 0xc000004078
// {娜扎 23}
// 0xc000004078
// {大王八 9000}

// 	fmt.Println(m)         //map[大王八:0xc000004078 娜扎:0xc000004078 小王子:0xc000004078]
// 	fmt.Printf("%#v\n", m) //map[string]*main.student{"大王八":(*main.student)(0xc000004078), "娜扎":(*main.student)(0xc000004078), "小王子":(*main.student)(0xc000004078)}

// 	for k, v := range m {
// 		fmt.Println(k, "=>", v.name)
// 	}
// }

//m是一个映射 ，make映射返回的是映射本身
//for _, stu := range stus ，遍历切片 ，返回的是结构体由stus接收
//m[小王子]=&stu
//m[娜扎]=&stu
//m[大王八]=&stu
//m:{"小王子":&main.student{name: "大王八", age: 9000} 娜扎:&main.student{name: "大王八", age: 9000} 大王八:&main.student{name: "大王八", age: 9000}}
// 小王子=>大王八
// 娜扎=>大王八
// 大王八=>大王八

package main

import "fmt"

func main() {
	type person struct {
		name string
		city string
		age  int8
	}

	p5 := person{ 
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
	fmt.Println(p5)            //{小王子 北京 18}
}
