package main

import "fmt"

type PaymentService struct {
	logger     *JSONLogger
	processors map[string]PaymentProcessing
}

func NewPaymentService(logger *JSONLogger) *PaymentService {
	return &PaymentService{
		logger:     logger,
		processors: make(map[string]PaymentProcessing),
	}
}

func (ps *PaymentService) RegisterProcessor(paymentID string, processor PaymentProcessing) {
	ps.processors[paymentID] = processor
}

func (ps *PaymentService) ProcessPayment(paymentID string, amount float64) error {
	processor, exists := ps.processors[paymentID]
	if !exists {
		return fmt.Errorf("payment processor not found for ID: %s", paymentID)
	}
	return processor.ProcessPayment(amount)
}

func (ps *PaymentService) Refund(paymentID string, amount float64) error {
	processor, exists := ps.processors[paymentID]
	if !exists {
		return fmt.Errorf("payment processor not found for ID: %s", paymentID)
	}
	return processor.Refund(amount)
}
