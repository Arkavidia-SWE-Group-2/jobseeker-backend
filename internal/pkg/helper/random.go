package helper

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		b[i] = letters[num.Int64()]
	}
	return string(b)
}

func VanityFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) < 2 {
		return ""
	}
	username := parts[0]

	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleanedUsername := re.ReplaceAllString(username, "")

	return cleanedUsername + RandomString(5)
}
