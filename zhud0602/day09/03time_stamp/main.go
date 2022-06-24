// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	// 获取北京时间所在的东八区时区对象
// 	secondsEastOfUTC := int((8 * time.Hour).Seconds())
// 	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

// 	// 北京时间 2022-02-22 22:22:22.000000022 +0800 CST
// 	t := time.Date(2022, 02, 22, 22, 22, 22, 22, beijing)
// 	fmt.Println(t)

// 	var (
// 		sec  = t.Unix()
// 		msec = t.UnixMilli()
// 		usec = t.UnixMicro()
// 	)

// 	// 将秒级时间戳转为时间对象（第二个参数为不足1秒的纳秒数）
// 	timeObj := time.Unix(sec, 22)
// 	fmt.Println(timeObj)           // 2022-02-22 22:22:22.000000022 +0800 CST
// 	timeObj = time.UnixMilli(msec) // 毫秒级时间戳转为时间对象
// 	fmt.Println(timeObj)           // 2022-02-22 22:22:22 +0800 CST
// 	timeObj = time.UnixMicro(usec) // 微秒级时间戳转为时间对象
// 	fmt.Println(timeObj)           // 2022-02-22 22:22:22 +0800 CST
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	fmt.Println(now)
// 	fmt.Println(now.Unix())
// 	fmt.Println(now.UnixNano())
// 	fmt.Println(now.Nanosecond())
// 	year, month, day := now.Date()
// 	fmt.Printf("year:%d month:%d day:%d\n", year, month, day)
// 	fmt.Println(now.Year())
// 	fmt.Println(now.Month())
// 	fmt.Println(now.Day())

// 	fmt.Println(now.Weekday())
// 	fmt.Println(now.YearDay())
// 	fmt.Println(now.Location())

// 	fmt.Println(now.Format("2006-01-02"))
// 	fmt.Println(now.Format("2006-01-02 15:04:05"))
// 	fmt.Println(now.Format("2006-01-02 03:04:05 AM"))

//time.Unix中time是包名，Unix是函数名。
//now.Unix()中now是time包中的Time类型，Unix()是接收体为Time类型的方法
//将时间戳返回Time类型，用函数time.Unix(sec int64, nsec int64)time.Time
// t := time.Unix(now.Unix(), 0)
// fmt.Println(t)
// layout := "2006-01-02 15:04:05"
// fmt.Println(t.Format(layout))

// layout := "2006-01-02 15:04:05"
// t := time.Date(2011, time.Month(3), 12, 15, 30, 20, 0, now.Location())
// fmt.Println(t.Format(layout))
// }

package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	// 1小时1分1s之后
	t1, _ := time.ParseDuration("1h1m1s")
	fmt.Println(t1)
	m1 := now.Add(t1)
	fmt.Println(m1)

	// 1小时1分1s之前
	t2, _ := time.ParseDuration("-1h1m1s")
	m2 := now.Add(t2)
	fmt.Println(m2)

	// 3小时之前
	t3, _ := time.ParseDuration("-1h")
	m3 := now.Add(t3 * 3)
	fmt.Println(m3)

	// 10 分钟之后
	t4, _ := time.ParseDuration("10m")
	m4 := now.Add(t4)
	fmt.Println(m4)

	// Sub 计算两个时间差
	sub1 := now.Sub(m3)
	fmt.Println(sub1.Hours())   // 相差小时数
	fmt.Println(sub1.Minutes()) // 相差分钟数
	fmt.Println(sub1)

	t5, _ := time.ParseDuration("2h")
	fmt.Println(t5)
	fmt.Println(now.Add(t5))
	fmt.Println(time.Now())
	fmt.Println(now.Add(1000000))
	fmt.Println(now.Format("2006/01/02 15:04:05.999"))

	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	timeObj, err := time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0000 UTC

	err = errors.New("无效的日志级别")
	fmt.Println(err)

}
