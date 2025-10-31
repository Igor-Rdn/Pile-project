package main

type PaymentProcessing interface {
	GetBalance() float64
	ProcessPayment(amount float64) error
	Refund(amount float64) error
}

type PaymentValidate interface {
	Validate() error
}

type PaymentLog interface {
	LogPayment(operation string, amount float64, success bool)
}
