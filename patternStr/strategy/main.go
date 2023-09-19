package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount int)
}

type CreditCardPayment struct{}

type KaspiQrPayment struct{}

func (cr *CreditCardPayment) Pay(amount int) {
	fmt.Printf("Payment in amount %v with card was successful \n", amount)
}

func (cr *KaspiQrPayment) Pay(amount int) {
	fmt.Printf("Payment in amount %v with Kaspi QR was successful\n"+
		"", amount)
}

type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func (sh *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	sh.paymentStrategy = strategy
}

func (sh *ShoppingCart) Checkout(amount int) {
	if sh.paymentStrategy == nil {
		fmt.Println("Please choose payment method")
		return
	}
	sh.paymentStrategy.Pay(amount)
}

func main() {
	creditCard := &CreditCardPayment{}
	kaspiQr := &KaspiQrPayment{}

	cart := &ShoppingCart{}

	cart.SetPaymentStrategy(creditCard)

	cart.Checkout(10000)

	cart.SetPaymentStrategy(kaspiQr)

	cart.Checkout(15000)
}
