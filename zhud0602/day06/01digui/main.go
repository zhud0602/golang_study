// package main

// import "fmt"

//递归：函数自己调用自己
//递归适合处理那种问题相同问题的规模越来越小的场景
//递归一定要有一个明确的退出条件
//计算n的阶乘
// func f(n uint) uint {
// 	if n <= 1 {
// 		return 1
// 	}
// 	return n * f(n-1)
// }

// func main() {
// 	fmt.Println(f(3)) //6
// }

//上台阶的面试题
//n个台阶，一次可以走1步，也可以走2步，有多少种走法。
package main

import "fmt"

func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	fmt.Println(taijie(4)) //5
}
