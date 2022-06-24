// package main

// import (
// 	"fmt"
// 	"strings"
// )

// /*
// 你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
// 分配规则如下：
// a. 名字中每包含1个'e'或'E'分1枚金币
// b. 名字中每包含1个'i'或'I'分2枚金币
// c. 名字中每包含1个'o'或'O'分3枚金币
// d: 名字中每包含1个'u'或'U'分4枚金币
// 写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
// 程序结构如下，请实现 ‘dispatchCoin’ 函数
// */
// var (
// 	coins = 50
// 	users = []string{
// 		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
// 	}
// 	distribution = make(map[string]int, len(users))
// )

// //1、遍历每个人名字
// //2、根据名字字符串拆分字符切片
// //3、遍历字符切片，根据字符规则求金币
// //3、根据人名字按照分金币规则求分到金币
// func dispatchCoin() int {
// 	for _, name := range users {
// 		arr := strings.Split(name, "") //名字字符切片
// 		sum := 0
// 		for _, word := range arr {
// 			switch word {
// 			case "e", "E":
// 				sum = sum + 1
// 			case "i", "I":
// 				sum = sum + 2
// 			case "o", "O":
// 				sum = sum + 3
// 			case "u", "U":
// 				sum = sum + 4
// 			}
// 			distribution[name] = sum
// 		}
// 		coins = coins - sum
// 	}
// 	fmt.Println(distribution)
// 	return coins
// }

// func main() {
// 	left := dispatchCoin()
// 	fmt.Println("剩下：", left)
// }

//map[Aaron:3 Adriano:5 Augustus:12 Elizabeth:4 Emilie:6 Giana:2 Heidi:5 Matthew:1 Peter:2 Sarah:0]
//剩下： 10

package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

//1、遍历每个人名字
//2、根据名字字符串拆分字符切片
//3、遍历字符切片，根据字符规则求金币
//3、根据人名字按照分金币规则求分到金币
func dispatchCoin() int {
	for _, name := range users {
		for _, word := range name {
			switch word {
			case 'e', 'E':
				distribution[name]++
			case 'i', 'I':
				distribution[name] += 2
			case 'o', 'O':
				distribution[name] += 3
			case 'u', 'U':
				distribution[name] += 4
			}
		}
		coins = coins - distribution[name]
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for index, value := range distribution {
		fmt.Printf("%s:%d\n", index, value)
	}

}

// 剩下： 10
// Giana:2
// Adriano:5
// Elizabeth:4
// Matthew:1
// Augustus:12
// Peter:2
// Aaron:3
// Heidi:5
// Emilie:6
