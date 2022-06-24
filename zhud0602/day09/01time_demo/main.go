package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	a, b, c := now.Date()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
	fmt.Println(a, b, c)
}
