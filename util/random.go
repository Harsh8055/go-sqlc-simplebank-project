package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

}

// generate a random string of length n
func RandomOwner() string {
	return StringWithCharset(5)
}

// generate a random string of length n
func StringWithCharset(length int) string {
	var sb strings.Builder

	for i := 0; i < length; i++{
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])  
	}
	return sb.String()
}

// generate a random int between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

/// generate a random int64 between 0 and 1000
func RandomMoney() int64 {
	return rand.Int63n(1000)
}

// generate a random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
