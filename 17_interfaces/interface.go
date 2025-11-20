package main

import "fmt"

type paymenter interface {
	//If you implement these methods, you can be treated as this interface
	//if there is same method signature in struct then it means
	// the struct is implementing interface implicitly

	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
	// razorpayPaymentGw.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {

	//logic to make payment
	fmt.Println("making payment using razor pay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

func main() {
	// stripePaymentGw := stripe{}
	razorpayPaymentGw := razorpay{}
	newPayment := payment{
		gateway: razorpayPaymentGw,
	}

	newPayment.makePayment(100)
	// stripePayment := stripe{}
	// stripePayment.pay(200)

}
