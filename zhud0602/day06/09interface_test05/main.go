// package main

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// type Closer interface {
// 	Close() error
// }

// // ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
// type ReadWriter interface {
// 	Reader
// 	Writer
// }

// // ReadCloser 是组合Reader接口和Closer接口形成的新接口类型
// type ReadCloser interface {
// 	Reader
// 	Closer
// }

// // WriteCloser 是组合Writer接口和Closer接口形成的新接口类型
// type WriteCloser interface {
// 	Writer
// 	Closer
// }

package main

import "fmt"

// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
		fmt.Println(v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	justifyType(100)
}
