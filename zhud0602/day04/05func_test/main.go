package main

import "fmt"

func main() {

	func(x, y int) (m int) {
		return x + y
	}(10, 20)

	fmt.Println() 
}
