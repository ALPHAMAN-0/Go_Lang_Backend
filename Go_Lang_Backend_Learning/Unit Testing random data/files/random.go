package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomName generates a random full name
func RandomName() string {
	firstNames := []string{"Alice", "Bob", "Charlie", "David", "Emma", "Frank", "Grace", "Henry"}
	lastNames  := []string{"Smith", "Johnson", "Brown", "Taylor", "Wilson", "Davis", "Clark"}
	return firstNames[rand.Intn(len(firstNames))] + " " + lastNames[rand.Intn(len(lastNames))]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@%s.com", RandomString(8), RandomString(5))
}

// RandomPhone generates a random phone number
func RandomPhone() string {
	return fmt.Sprintf("+1-%03d-%03d-%04d",
		RandomInt(100, 999),
		RandomInt(100, 999),
		RandomInt(1000, 9999),
	)
}
