package main

import "fmt"

//拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
// func main() {
// 	a1 := [3]int{1, 2, 3}
// 	b1 := a1[1:]
// 	b2 := b1
// 	b2[0] = 100
// 	fmt.Println(b1, b2)
// }

//[100 3] [100 3]

//声明切片类型
// var a1 []string
// var a2 []int
// var a3 []bool

// //初始化切片
// var b1 = []int{}
// var b2 = []bool{}

// func main() {

// 	fmt.Println(b1, b2)
// 	fmt.Println(a1 == nil) //true
// 	fmt.Println(b1 == nil)  //false
// 	//fmt.Println(b1==b2) //切片是引用类型，不支持直接比较，只能和nil比较
// }

// func main() {
// 	a := [5]int{1, 2, 3, 4, 5}
// 	s := a[1:3]
// 	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
// }

//s:[2 3] len(s):2 cap(s):4

// func main() {
// 	var s1 []int
// 	s2 := []int{}
// 	s3 := make([]int, 0)
// 	fmt.Println(s1 == nil) //true
// 	fmt.Println(s2 == nil) //false
// 	fmt.Println(s3 == nil) //false
// 	fmt.Printf("%v\n", s2)
// 	fmt.Println(s1 == nil)
// }

// 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil

// func main() {
// 	s1 := []int{1, 2, 3, 4, 5}
// 	s1 = append(s1, 6)
// 	fmt.Println(s1)
// 	fmt.Printf("len(s1):%v,cap(s1):%v\n", len(s1), cap(s1))

// }

//调用append函数必须用原来的切片变量来接收返回值
//append追加元素，原来的底层数组放不下的时候，GO底层就会把底层数组换一个
//[1 2 3 4 5 6]

// func main() {
// 	a1 := []int{1, 3, 5, 7, 9}
// 	a2 := a1
// 	a3 := make([]int, 5, 5)
// 	copy(a3, a1)  ////使用copy()函数将切片a中的元素复制到切片c
// 	a1[0] = 100
// 	fmt.Println(a1)
// 	fmt.Println(a2)
// 	fmt.Println(a3)
// [100 3 5 7 9]
// [100 3 5 7 9]
// [1 3 5 7 9]
//}

// func main() {
// 	a1 := []int{1, 3, 5, 7, 9}
// 	a2 := a1
// 	a3 := make([]int, 5, 5)
// 	a1[0] = 100
// 	copy(a3, a1)
// 	fmt.Println(a1)
// 	fmt.Println(a2)
// 	fmt.Println(a3)
// [100 3 5 7 9]
// [100 3 5 7 9]
// [100 3 5 7 9]
// }

// func main() {
// 	var s []int
// 	s = append(s, 1)
// 	s = append(s, 2, 3, 4)
// 	s2 := []int{5, 6, 7}
// 	s = append(s, s2...)

// 	fmt.Println(s)  //[1 2 3 4 5 6 7]

// }

//Go语言的内建函数append()可以为切片动态添加元素。 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。

// func main() {
// 	var numslice []int
// 	for i := 0; i < 10; i++ {
// 		numslice = append(numslice, i)
// 		fmt.Printf("%v  len:%d   cap:%d  ptr:%p\n", numslice, len(numslice), cap(numslice), numslice)
// 	}
// }

// [0 1]  len:2   cap:2  ptr:0xc0000160e0
// [0 1 2]  len:3   cap:4  ptr:0xc000014220
// [0 1 2 3]  len:4   cap:4  ptr:0xc000014220
// [0 1 2 3 4]  len:5   cap:8  ptr:0xc0000122c0
// [0 1 2 3 4 5]  len:6   cap:8  ptr:0xc0000122c0
// [0 1 2 3 4 5 6]  len:7   cap:8  ptr:0xc0000122c0
// [0 1 2 3 4 5 6 7]  len:8   cap:8  ptr:0xc0000122c0
// [0 1 2 3 4 5 6 7 8]  len:9   cap:16  ptr:0xc000010200
// [0 1 2 3 4 5 6 7 8 9]  len:10   cap:16  ptr:0xc000010200

// func main() {
// 	var s = []int{}
// 	for i := 1; i < 10; i++ {
// 		s = append(s, i)
// 		fmt.Printf("%v  len(s):%d  cap(s):%d\n", s, len(s), cap(s))
// 	}

// }

// func main() {
// 	var s = []int{1, 2, 3}
// 	s = append(s, 4, 5, 6, 7, 8, 9, 10, 11)
// 	fmt.Printf("%v  len(s):%d  cap(s):%d\n", s, len(s), cap(s))
// }

//[1 2 3 4 5 6 7 8 9 10 11]  len(s):11  cap(s):12
//int类型切片扩容是+2，+2的扩容

// func main() {
// 	var s = []string{"北京", "上海", "广州"}
// 	s = append(s, "武汉", "西安", "杭州", "深圳", "合肥", "石家庄")

// 	//将s中的"西安"，"合肥"删掉
// 	s = append(s[0:4], s[5:]...)
// 	fmt.Printf("%v  len(s1):%d  cap(s1):%d\n", s, len(s), cap(s))
// 	//[北京 上海 广州 武汉 杭州 深圳 合肥 石家庄]  len(s1):8  cap(s1):9

// 	s1 := s
// 	fmt.Println(s1)
// 	// [北京 上海 广州 武汉 杭州 深圳 合肥 石家庄]
// 	s1 = append(s1[:6], s1[7:]...)
// 	fmt.Printf("%v  len(s1):%d  cap(s1):%d\n", s1, len(s1), cap(s1))
// 	// [北京 上海 广州 武汉 杭州 深圳 石家庄]  len(s1):7  cap(s1):9
// }

//[北京 上海 广州 武汉 西安 杭州 深圳 合肥 石家庄]  len(s):9  cap(s):9
//string类型切片扩容是+1,+1的扩容

// func main() {
// 	var s = [3]string{"北京", "上海", "广州"}
// 	s1 := s[:]
// 	fmt.Printf("s1:%v  len(s1):%d  cap(s1):%d\n", s1, len(s1), cap(s1))

// 	s1 = append(s1[0:1], s1[2:3]...)

// 	fmt.Printf("s:%v  len(s):%d  cap(s):%d\n", s, len(s), cap(s))
// 	fmt.Printf("s1:%v  len(s1):%d  cap(s1):%d\n", s1, len(s1), cap(s1))

// }

//  s1:[北京 上海 广州]  len(s1):3  cap(s1):3
//  s:[北京 广州 广州]  len(s):3  cap(s):3
//  s1:[北京 广州]  len(s1):2  cap(s1):3
//可能解释： 切片s1在原来数组s基础上进行元素删除后，底层数组仍然是s，这时候切片中元素的删除会影响原数组s

// func main() {
// 	var s = [3]string{"北京", "上海", "广州"}
// 	s1 := s[:]
// 	s1 = append(s1, "武汉", "西安", "杭州", "深圳", "合肥", "石家庄")
// 	fmt.Printf("s1:%v  len(s1):%d  cap(s1):%d\n", s1, len(s1), cap(s1))
// 	fmt.Printf("s:%v  len(s):%d  cap(s):%d\n", s, len(s), cap(s))
// 	s1 = append(s1[0:1], s1[2:3]...)
// 	fmt.Printf("s1:%v  len(s1):%d  cap(s1):%d\n", s1, len(s1), cap(s1))
// 	//fmt.Printf("s1:%T  s1:%+v\n", s1, s1)
// 	fmt.Printf("s:%v  len(s):%d  cap(s):%d\n", s, len(s), cap(s))
// 	//fmt.Printf("s:%T", s)
// }

// s1:[北京 上海 广州 武汉 西安 杭州 深圳 合肥 石家庄]  len(s1):9  cap(s1):9
// s:[北京 上海 广州]  len(s):3  cap(s):3
// s1:[北京 广州]  len(s1):2  cap(s1):9
// s:[北京 上海 广州]  len(s):3  cap(s):3
//可能解释： 切片在原来数组s基础上扩容后，底层数组不再是s，这时候切片中元素的删除不会影响原数组s

// func main() {
// 	arr := []int{1, 2, 3}
// 	//newArr := []*int{}
// 	for i, _ := range arr {
// 		fmt.Println(arr[i])
// 		fmt.Println(&arr[i])
// 		//newArr = append(newArr, &v)
// 	}
// 	// for _, v := range newArr {
// 	// 	fmt.Println(*v)
// 	// }
// }

func main() {
	var out []*int
	for i := 0; i < 5; i++ {
		m := &i
		fmt.Println(m)
		out = append(out, &i)
	}
	//fmt.Println("Values:", *out[0])
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	//fmt.Println("Addresses:", out[0], out[1], out[2])
}
