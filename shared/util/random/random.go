package random

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func Int(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func String(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func Name() string {
	return String(6)
}

func PhoneNumber() string {
	areaCode := [5]string{"0412", "0414", "0416", "0424", "0426"}
	randomAreaCode := areaCode[Int(0, 4)]
	randomNumber := fmt.Sprintf("%07d", rand.Intn(10000000))

	return fmt.Sprintf("%s%s", randomAreaCode, randomNumber)
}

func IndentificationNumber() string {
	randomNumber := fmt.Sprintf("%09d", rand.Intn(1000000000))
	return randomNumber
}

func IdentificationType() string {
	it := [5]string{"V", "E", "P", "J", "G"}
	return it[Int(0, 4)]
}

func Email() string {
	return fmt.Sprintf("%s@email.com", String(6))
}
