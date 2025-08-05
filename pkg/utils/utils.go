package utils

import (
	"math/rand"
	"time"
)

func GenerateCardNumber() string {
	rand.Seed(time.Now().UnixNano())
	const digits = "0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = digits[rand.Intn(len(digits))]
	}
	return string(b)
}

func GenerateReferenceID() string {
	rand.Seed(time.Now().UnixNano())
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 12)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
