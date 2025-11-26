package main

import "fmt"


type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed 
	Prepared 
	Delivered
)


func changeOrderStatus(status string) {
	fmt.Println("chaning order status to" , status)
}
func main() {
	changeOrderStatus("confirmed")
}
