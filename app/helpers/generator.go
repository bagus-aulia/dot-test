package helpers

import (
	"math/rand"
	"time"
)

// GenerateUUID id function to generate UUID by it's param
func GenerateUUID(param string) string {
	uuid := RandString(5) + "-" + RandString(7) + "-" + RandString(4) + param[0:1] + RandString(3) + "-" + time.Now().Format("060102") + time.Now().Format("150405")

	return uuid
}

// RandString is function to generate random string
func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
