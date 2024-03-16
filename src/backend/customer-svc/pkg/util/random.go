package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return RandomString(6)
}

func RandomPhoneNumber() string {
	areaCode := [5]string{"0412", "0414", "0416", "0424", "0426"}
	randomAreaCode := areaCode[RandomInt(0, 4)]
	randomNumber := fmt.Sprintf("%07d", rand.Intn(10000000))

	return fmt.Sprintf("%s%s", randomAreaCode, randomNumber)
}

func RandomIndentificationNumber() string {
	randomNumber := fmt.Sprintf("%09d", rand.Intn(1000000000))
	return randomNumber
}

func RandomIdentificationType() string {
	it := [5]string{"V", "E", "P", "J", "G"}
	return it[RandomInt(0, 4)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
