package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota   //use "" for string
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}
func main() {
	changeOrderStatus(Received)
}
