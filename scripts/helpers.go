package payment_processor

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
)

// ValidateCardNumber checks if a card number is valid
func ValidateCardNumber(cardNumber string) bool {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	cardNumber = strings.ReplaceAll(cardNumber, "-", "")

	if len(cardNumber) < 13 || len(cardNumber) > 19 {
		return false
	}

	sum := 0
	for i := range cardNumber {
		digit, _ := strconv.Atoi(string(cardNumber[i]))
		if i%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}

// ValidateCVV checks if a CVV is valid
func ValidateCVV(cvv string) bool {
	_, err := strconv.Atoi(cvv)
	if err != nil {
		return false
	}

	return len(cvv) == 3 || len(cvv) == 4
}

// GetCardExpirationDate checks if a card expiration date is valid
func GetCardExpirationDate(cardExpirationDate string) (int, int, error) {
	parts := strings.Split(cardExpirationDate, "/")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid expiration date format")
	}

	month, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	if month < 1 || month > 12 {
		return 0, 0, fmt.Errorf("invalid month")
	}

	year, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	now := time.Now()
	if year < now.Year() || (year == now.Year() && month < now.Month()) {
		return 0, 0, fmt.Errorf("card has expired")
	}

	return month, year, nil
}

// ProcessPayment processes a payment
func ProcessPayment(payment Payment) (*PaymentConfirmation, error) {
	producer := sarama.NewSyncProducer([]string{"localhost:9092"})

	msg := &sarama.ProducerMessage{
		Topic: "payments",
		Value:  sarama.StringEncoder(payment),
	}

	producer.Send(msg)
	producer.Close()

	return &PaymentConfirmation{
		ID:        payment.ID,
		Status:    "pending",
		CreatedAt: time.Now(),
	}, nil
}

// SaveToRedis saves a payment to Redis
func SaveToRedis(payment Payment) error {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := client.Set(context.Background(), "payment:"+strconv.Itoa(payment.ID), payment, 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}