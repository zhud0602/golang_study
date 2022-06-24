package main

import "fmt"

type Zhifubao struct{}

func (z *Zhifubao) Pay(amount int64) {
	fmt.Printf("使用支付宝付款:%.2f元\n", float64(amount/100))
}

// func Checkout(obj *Zhifubao) {
// 	obj.Pay(100)
// }

type Wechat struct{}

func (w *Wechat) Pay(amount int64) {
	fmt.Printf("使用微信付款:%.2f元\n", float64(amount/100))
}

type Payer interface {
	Pay(int64)
}

func Checkout(obj Payer) {
	obj.Pay(100)
}

func main() {
	Checkout(&Zhifubao{})
	Checkout(&Wechat{})
}
