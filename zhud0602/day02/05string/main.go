package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//s := "hello 沙河"
	//字符串，用双引号包裹
	//c1 := 'H'
	//C2 := '1'
	//C3 := '沙'
	//单独的字母、汉字、符号表示一个字符，用单引号包裹

	fmt.Println("str := \"D:\\PA\\workspace\"")   //str := "D:\PA\workspace"    \"  含义双引号
	fmt.Println("str := \"D:\\PA\\work\rspace\"") //space" "D:\PA\work          \\  含义反斜杠      \r含义 返回行首
	fmt.Println("str := \"D:\\PA\\work\nspace\"") //    \n 含义换行符（直接跳到下一行的同列位置）
	//str := "D:\PA\work
	//space"
	fmt.Println("str := \"D:\tPA\tworkspace\"") //    \t 含义制表符   str := "D:      PA      workspace"
	fmt.Println("str := 'D:\\PA\\workspace'")   //str := 'D:\PA\workspace'  单引号 可以直接用'表示
	fmt.Println("str := 'D:\\PA\\workspace'")   //

	s1 := `第一行
	第二行    
	第三行
			第四行
	`
	fmt.Println(s1)
	/*第一行
	第二行
	第三行
				第四行*/

	tip1 := "genji is a ninja"
	fmt.Println(len(tip1)) //16
	tip2 := "忍者"
	fmt.Println(len(tip2)) //6    Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节，因此使用 len() 获得两个中文文字对应的 6 个字节。

	fmt.Println(utf8.RuneCountInString(tip1)) //16   ASCII 字符串长度使用 len() 函数。
	fmt.Println(utf8.RuneCountInString(tip2)) //2    Unicode 字符串长度使用 utf8.RuneCountInString() 函数。

	s2 := tip1 + tip2 //拼接字符串 +  ，或者用fmt.Sprintf
	fmt.Println(s2)

	s3 := fmt.Sprintf("%s is %s", tip1, tip2)
	fmt.Println(s3) //genji is a ninja is 忍者

	fmt.Println(strings.Split(tip1, " "))         //[genji is a ninja]  //strings.Split 分割字符串
	fmt.Println(strings.Contains(tip1, "ninja"))  //true ,  strings.contains 判断是否包含
	fmt.Println(strings.HasPrefix(tip1, "genji")) //true , strings.HasPrefix 前缀判断
	fmt.Println(strings.HasSuffix(s3, "忍者"))      //true , strings.HasSuffix 后缀判断
	fmt.Println(strings.Index(s3, "is"))          //6  子串出现的位置
	fmt.Println(strings.LastIndex(s3, "is"))      //17  子串最后一次出现的位置

	// s4 := strings.Split(tip1, " ")
	// fmt.Println(strings.Join(s4, "-"))
	/*func Join(s []string, sep string) string
	  在这里，s是可用来连接元素的字符串，sep是放置在最终字符串中元素之间的分隔符。*/

	n1 := 32
	fmt.Printf("value:%v,type:%T\n", n1, n1)

	n2 := 1.2345
	fmt.Printf("value:%v,type:%T\n", n2, n2)

	n3 := false
	fmt.Printf("value:%v,type:%T\n", n3, n3)

	n4 := "teacher"
	fmt.Printf("value:%v,type:%T\n", n4, n4)

}
