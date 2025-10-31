package main

import "time"

type Payment struct {
	PaymentID   string
	Amount      float64
	Currency    string
	CreatedAt   time.Time
	PaymentType int
}

// PaymentType - тип платежной системы
type PaymentType int

const (
	Card PaymentType = iota + 1
	PayPal
	Crypto
)

//PaymentType = 1
type CardPayment struct {
	Payment
	CardNumber     string
	ExpiryDate     time.Time
	CVV            string
	CurrentBalance float64
	Logger         *JSONLogger
}

//PaymentType = 2
type PayPalPayment struct {
	Payment
	Email          string
	Password       string
	CurrentBalance float64
}

//PaymentType = 3
type CryptoPayment struct {
	Payment
	WalletAddress  string
	PrivateKey     string
	CurrentBalance float64
}
