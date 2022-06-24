package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!,error:", err)
		return
	}

	defer file.Close()

	var content []byte
	var temp = make([]byte, 128)
	for {
		n, err := file.Read(temp)
		if err == io.EOF {
			fmt.Println("文件读取完了")
			break

		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
		}
		content = append(content, temp[:n]...)
	}
	fmt.Println(string(content))
}
