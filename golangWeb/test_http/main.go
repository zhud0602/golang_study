package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func testGet() {
	url1 := "https://www.baidu.com"
	resp, err := http.Get(url1)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	readAll, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("b:%v\n", string(readAll))
}

func main() {
	testGet()
}
