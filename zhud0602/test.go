//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	for i := 0; i <= 10; i++ {
//		go func() {
//			fmt.Println(i)
//		}()
//	}
//	fmt.Println("main")
//	time.Sleep(time.Second)
//}

//7
//9
//11
//11
//main
//11
//11
//11
//11
//11
//11
//11

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
////程序启动之后会创建一个主goroutine去执行
//func main() {
//	for i := 0; i <= 10; i++ {
//		//开启一个单独的goroutine去执行hello函数(任务)
//		go func(i int) {
//			fmt.Println(i) //用的是函数参数的i，不是函数外面的i
//		}(i)
//	}
//	fmt.Println("main")
//	time.Sleep(time.Second)
//	//main函数结束了，由main函数启动的goroutine也结束了
//}

//main
//1
//6
//0
//3
//5
//8
//7
//9
//4
//10
//2

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func f() {
//	rand.Seed(time.Now().UnixNano())
//	for i := 0; i < 5; i++ {
//		n1 := rand.Int()
//		n2 := rand.Intn(10)
//		fmt.Println(-n1, -n2)
//	}
//}
//
//func main() {
//	f()
//
//}

//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func main() {
//	fmt.Println(runtime.NumCPU())
//
//}

//package main
//
//import (
//	"fmt"
//	"os"
//	"runtime/trace"
//)
//
//func main() {
//
//	//创建trace文件
//	f, err := os.Create("trace.out")
//	if err != nil {
//		panic(err)
//	}
//
//	defer f.Close()
//
//	//启动trace goroutine
//	err = trace.Start(f)
//	if err != nil {
//		panic(err)
//	}
//	defer trace.Stop()
//
//	//main
//	fmt.Println("Hello World")
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//var wg sync.WaitGroup
//
//func hello(i int) {
//	fmt.Println("hello", i)
//	wg.Done()
//}
//
//func main() {
//
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go hello(i)
//	}
//
//	fmt.Println("你好")
//	wg.Wait()
//}
//package main
//
//import "fmt"
//
//func main() {
//
//	for i := 0; i < 5; i++ {
//		go func() {
//			fmt.Println(i)
//		}()
//	}
//
//}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//var ch1 = make(chan int, 50)
//var ch2 = make(chan int, 100)
//var wg sync.WaitGroup
//
//func main() {
//
//	wg.Add(2)
//
//	go func() {
//		defer wg.Done()
//		rand.Seed(time.Now().Unix())
//		for i := 0; i < 100; i++ {
//			ch1 <- i
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		//for i := 0; i < 100; i++ {
//		//	x := <-ch1
//		//	//ch2 <- math.Pow(2, float64(x))
//		//	ch2 <- x * x
//		for m := range ch1 {
//			ch2 <- m * m
//		}
//	}()
//
//	//for x := range ch2 {
//	//	fmt.Println(x)
//	//}
//
//	for i := 0; i < 100; i++ {
//		x := <-ch2
//		fmt.Println(x)
//	}
//
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//var wg sync.WaitGroup
//
//func f1(ch1 chan int) {
//	defer wg.Done()
//	for i := 0; i < 100; i++ {
//		ch1 <- i
//	}
//}
//
//func f2(ch1, ch2 chan int) {
//	defer wg.Done()
//	for x := range ch1 {
//		ch2 <- x * x
//	}
//}
//
//func main() {
//	a := make(chan int, 100)
//	b := make(chan int, 100)
//	wg.Add(2)
//	go f2(a, b)
//	go f1(a)
//	wg.Wait()
//	for ret := range b {
//		fmt.Println(ret)
//	}
//}

package main

import (
	"fmt"
	"sync"
)

// demo1 通道误用导致的bug

func main() {
	wg := sync.WaitGroup{}
	//once := sync.Once{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
		flo:
			for {
				//task := <-ch
				// 这里假设对接收的数据执行某些操作

				//task, ok := <-ch
				//if ok {
				//	fmt.Println(task)
				//} else {
				//	break
				//}

				//select {
				//case task, ok := <-ch:
				//	if !ok {
				//		ch = nil
				//	}
				//	fmt.Println(task)
				//default:
				//	break flo
				//}

				select {
				case task, ok := <-ch:
					if !ok {
						break flo
					}
					fmt.Println(task)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
