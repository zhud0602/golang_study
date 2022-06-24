package main

import "fmt"

// func main() {
// BREAKDEMO1:
// 	for i := 0; i < 10; i++ {
// 		for j := 'a'; j < 'z'; j++ {
// 			if j == 'b' {
// 				break BREAKDEMO1
// 			}
// 			fmt.Printf("%v-%c\n", i, j)
// 		}
// 	}
// 	fmt.Println("over")
// }

// // 0-a
// // over

// func main() {
// 	for i := 0; i < 10; i++ {
// 		for j := 'a'; j < 'd'; j++ {
// 			if j == 'b' {
// 				break
// 			}
// 			fmt.Printf("%v-%c\n", i, j)
// 		}
// 	}
// 	fmt.Println("over")
// }

/*
0-a
1-a
2-a
3-a
4-a
5-a
6-a
7-a
8-a
9-a
over
*/

func main() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

/*
0-0
0-1
0-2
0-3
0-4
1-0
1-1
1-2
1-3
1-4
2-0
2-1
3-0
3-1
3-2
3-3
3-4
4-0
4-1
4-2
4-3
4-4
*/
