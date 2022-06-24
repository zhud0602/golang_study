package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timestamp := now.Unix()
	milli := now.UnixMilli()
	micro := now.UnixMicro()
	nano := now.UnixNano()
	xx := now.Local().Unix()
	fmt.Println(timestamp, milli, micro, nano, xx)
}
