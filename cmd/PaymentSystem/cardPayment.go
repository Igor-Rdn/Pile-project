package main

import (
	"fmt"
	"time"
)

func NewCardPayment(cardNumber string, expiry time.Time, balance float64, currency string, cvv string, logger *JSONLogger) *CardPayment {

	res := CardPayment{
		Payment: Payment{
			PaymentID:   GenerateID(),
			PaymentType: 1,
			CreatedAt:   time.Now(),
			Currency:    currency,
		},
		CardNumber:     cardNumber,
		ExpiryDate:     expiry,
		CurrentBalance: balance,
		CVV:            cvv,
		Logger:         logger,
	}

	return &res

}

// Проверка баланса
func (c *CardPayment) GetBalance() float64 {
	return c.CurrentBalance
}

// Списание средств
func (c *CardPayment) ProcessPayment(amount float64) error {

	if err := c.validate(); err != nil {
		c.Logger.LogPayment("payment", "card", amount, false, c.PaymentID, HideCardNumber(c.CardNumber), err)
		return err
	}

	if amount <= 0 {
		err := fmt.Errorf("incorrect amount: %.2f", amount)
		c.Logger.LogPayment("payment", "Card", amount, false, c.PaymentID, HideCardNumber(c.CardNumber), err)
		return err
	}

	if c.GetBalance() < amount {
		err := fmt.Errorf("insufficient funds: have %.2f, need %.2f", c.CurrentBalance, amount)
		c.Logger.LogPayment("payment", "Card", amount, false, c.PaymentID, HideCardNumber(c.CardNumber), err)
		return err
	}

	c.CurrentBalance -= amount
	c.Logger.LogPayment("payment", "Card", amount, true, c.PaymentID, HideCardNumber(c.CardNumber), nil)
	return nil
}

// Возврат средств
func (c *CardPayment) Refund(amount float64) error {

	if amount <= 0 {
		err := fmt.Errorf("incorrect refund amount: %.2f", amount)
		c.Logger.LogPayment("refund", "Card", amount, false, c.PaymentID, HideCardNumber(c.CardNumber), err)
		return err
	}

	c.CurrentBalance += amount
	c.Logger.LogPayment("refund", "Card", amount, true, c.PaymentID, HideCardNumber(c.CardNumber), nil)
	return nil
}

func (c *CardPayment) validate() error {

	if len(c.CardNumber) != 16 {
		return fmt.Errorf("invalid CardNumber")
	}

	if len(c.CVV) != 3 {
		return fmt.Errorf("incorrect CVV")
	}

	if time.Now().After(c.ExpiryDate) {
		return fmt.Errorf("card expired")
	}

	return nil
}
