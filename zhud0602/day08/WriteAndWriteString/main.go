package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河\n"
	file.Write([]byte(str))
	file.WriteString("hello 小王子")
}
