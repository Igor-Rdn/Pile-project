package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Инициализация логгера
	logger, err := NewJSONLogger("payments.log")
	if err != nil {
		log.Fatal("Failed to create logger:", err)
	}
	defer logger.Close()

	// Создание платежного сервиса
	paymentService := NewPaymentService(logger)

	// Создание карточного платежа
	cardPayment := NewCardPayment(
		"1234567812345678",
		time.Now().AddDate(2, 0, 0), // карта действует 2 года
		1000.0,
		"USD",
		"123",
		logger,
	)

	// Регистрация процессора в сервисе
	paymentService.RegisterProcessor(cardPayment.PaymentID, cardPayment)

	fmt.Println("=== Payment System Test ===")
	fmt.Printf("Initial balance: $%.2f\n", cardPayment.GetBalance())

	// Тест 1: Успешный платеж
	fmt.Println("\n1. Processing payment of $500...")
	if err := paymentService.ProcessPayment(cardPayment.PaymentID, 500.0); err != nil {
		fmt.Printf("Payment failed: %v\n", err)
	} else {
		fmt.Printf("Payment successful! New balance: $%.2f\n", cardPayment.GetBalance())
	}

	// Тест 2: Недостаточно средств
	fmt.Println("\n2. Processing payment of $600...")
	if err := paymentService.ProcessPayment(cardPayment.PaymentID, 600.0); err != nil {
		fmt.Printf("Payment failed: %v\n", err)
	}

	// Тест 3: Возврат средств
	fmt.Println("\n3. Processing refund of $200...")
	if err := paymentService.Refund(cardPayment.PaymentID, 200.0); err != nil {
		fmt.Printf("Refund failed: %v\n", err)
	} else {
		fmt.Printf("Refund successful! New balance: $%.2f\n", cardPayment.GetBalance())
	}

	// Тест 4: Неверная сумма
	fmt.Println("\n4. Processing payment of $-100...")
	if err := paymentService.ProcessPayment(cardPayment.PaymentID, -100.0); err != nil {
		fmt.Printf("Payment failed: %v\n", err)
	}

	fmt.Println("\n=== Test completed ===")
	fmt.Println("Check payments.log for detailed transaction log")
}
