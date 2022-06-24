package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//将一个文本文件拷贝到另外一个文件中

func main() {

	//1、打开source文件并读取
	filePath := "d:/source.txt"
	buf, err := ioutil.ReadFile(filePath)
	//fp2, err2 := os.Open("D:/target.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fp, err2 := os.Open("d:/target.txt")
	fp, err := os.OpenFile("d:/target.txt", os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		err = fp.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err3 := fp.Write(buf)
	if err3 != nil {
		fmt.Println(err3)
	}
}
