package main

import (
	"crypto/rand"
	"fmt"
)

func HideCardNumber(cardNumber string) (newCardNumber string) {

	newCardNumber = "**** **** **** " + cardNumber[12:]
	return
}

func GenerateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
