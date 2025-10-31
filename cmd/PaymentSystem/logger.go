package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp        time.Time `json:"timestamp"`
	Level            string    `json:"level"`       // info, error
	Operation        string    `json:"operation"`   // payment, refund, validate
	PaymentType      string    `json:"paymentType"` // card, paypal, crypto
	Amount           float64   `json:"amount"`
	Status           string    `json:"status"` // success, failed
	Error            string    `json:"error,omitempty"`
	PaymentID        string    `json:"paymentID"`
	HiddenCardNumber string    `json:"HiddenCardNumber,omitempty"`
}

type JSONLogger struct {
	file *os.File
	mu   sync.Mutex
}

func NewJSONLogger(filename string) (*JSONLogger, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	return &JSONLogger{file: file}, nil
}

func (l *JSONLogger) LogPayment(operation, paymentType string, amount float64, success bool, paymentID string, hiddenCardNumber string, err error) {

	entry := LogEntry{
		Timestamp:   time.Now(),
		Operation:   operation,
		PaymentType: paymentType,
		Amount:      amount,
		PaymentID:   paymentID,
		Status:      "SUCCESS",
		Level:       "INFO",
	}

	if !success {
		entry.Status = "FAILED"
		entry.Level = "ERROR"
		if err != nil {
			entry.Error = err.Error()
		}
	}

	if hiddenCardNumber != "" {
		entry.HiddenCardNumber = hiddenCardNumber
	}

	l.writeLogEntry(entry)
}

func (l *JSONLogger) writeLogEntry(entry LogEntry) {
	l.mu.Lock()
	defer l.mu.Unlock()

	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Printf("Failed to marshal log entry: %v\n", err)
		return
	}

	data = append(data, '\n')

	if _, err := l.file.Write(data); err != nil {
		fmt.Printf("Failed to write log: %v\n", err)
	}
}

func (l *JSONLogger) Close() error {
	return l.file.Close()
}
