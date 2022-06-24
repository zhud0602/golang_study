package main

import "fmt"

// func main() {
// 	scoremap := make(map[string]int, 8)
// 	scoremap["小明"] = 90
// 	scoremap["张三"] = 60
// 	fmt.Println(scoremap)
// 	fmt.Println(scoremap["小明"])
// 	fmt.Printf("type of scoremap:%T\n", scoremap)

// }

// map[小明:90 张三:60]
// 90
// type of scoremap:map[string]int

// func main() {
// 	userinfo := map[string]string{
// 		"username": "张三",
// 		"password": "123456",
// 	}
// 	fmt.Println(userinfo)
// }
//map[password:123456 username:张三]

// func main() {
// 	scoremap := map[string]string{
// 		"张三": "80",
// 		"李四": "90",
// 	}

// 	//value, bool := map[key]  value接收map[key]的值，bool判断是否存在key的键

// 	value, bool := scoremap["张三"]
// 	if bool {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("查无此人")
// 	}

// }

//80

//map的for range 遍历

// func main() {
// 	scoremap := map[string]int{
// 		"张三": 90,
// 		"李四": 80,
// 		"王五": 70,
// 	}

// 	for key, v := range scoremap {
// 		fmt.Printf("string:%s int:%d\n", key, v)
// 	}

// }

// string:张三 int:90
// string:李四 int:80
// string:王五 int:70

// func main() {
// 	scoreMap := make(map[string]int)
// 	scoreMap["张三"] = 90
// 	scoreMap["小明"] = 100
// 	scoreMap["娜扎"] = 60
// 	for _, value := range scoreMap {
// 		fmt.Println(value)
// 	}
// }

// 张三
// 小明
// 娜扎

//删除键值对，格式如下：delete(map,key)
//map:表示要删除键值对的名称   key：表示要删除的键值对的键
// func main() {
// 	scoreMap := make(map[string]int)
// 	scoreMap["张三"] = 100
// 	scoreMap["李四"] = 80
// 	scoreMap["王五"] = 70
// 	delete(scoreMap, "王五")
// 	for key,value:=range scoreMap{
// 		fmt.Println(key,value)
// 	}
// }

// 张三 100
// 李四 80

//值为切片类型的map
// func main() {
// 	var sliceMap = make(map[string][]string, 3)
// 	fmt.Println(sliceMap)     //map[]
// 	fmt.Println("after init") //after init
// 	key := "中国"
// 	value, ok := sliceMap[key]
// 	if !ok {
// 		value = make([]string, 0, 2)
// 	}
// 	value = append(value, "北京", "上海")
// 	sliceMap[key] = value
// 	fmt.Println(sliceMap) //map[中国:[北京,上海]]
// }

//元素为map类型的切片
// func main() {
// 	var mapSlice = make([]map[string]string, 3)
// 	for index, value := range mapSlice {
// 		fmt.Printf("index:%d value:%v\n", index, value) //0, 1, 2,
// 	}
// 	fmt.Println("after init")
// 	// 对切片中的map元素进行初始化
// 	mapSlice[0] = make(map[string]string, 10)
// 	mapSlice[0]["name"] = "小王子"
// 	mapSlice[0]["password"] = "123456"
// 	mapSlice[0]["address"] = "沙河"
// 	for index, value := range mapSlice {
// 		fmt.Printf("index:%d value:%v\n", index, value)
// 	}
// }

// index:0 value:map[]
// index:1 value:map[]
// index:2 value:map[]
// after init
// index:0 value:map[address:沙河 name:小王子 password:123456]
// index:1 value:map[]
// index:2 value:map[]

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)         //[1 2 3]
	m["q1mi"] = s                  //map[qimi:s]
	s = append(s[:1], s[2:]...)    //s=[]int{1,3}
	fmt.Printf("%+v\n", s)         //[1 3]
	fmt.Printf("%+v\n", m["q1mi"]) //[1 3 3]
}
