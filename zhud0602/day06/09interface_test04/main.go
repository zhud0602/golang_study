package main

import "fmt"

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct{}

func (d dryer) dry() {
	fmt.Println("甩一甩")
}

type haier struct {
	dryer
}

func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	var w WashingMachine
	h1 := &haier{}
	w = h1
	w.wash()
	w.dry()
}
