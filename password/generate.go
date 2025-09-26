package password

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	symbolChars  = "!@#$%^&*()-_=+[]{}<>?/|"
	defaultChars = lowerChars + upperChars + numberChars
)

// Generate membuat password acak
func Generate(length int, includeSymbols, includeNumbers, includeUppercase bool) (string, error) {
	var password strings.Builder
	chars := lowerChars

	if includeUppercase {
		chars += upperChars
	}
	if includeNumbers {
		chars += numberChars
	}
	if includeSymbols {
		chars += symbolChars
	}

	for i := 0; i < length; i++ {
		randomChar, err := getRandom(chars)
		if err != nil {
			return "", err
		}
		password.WriteByte(randomChar)
	}

	return password.String(), nil
}

func getRandom(chars string) (byte, error) {
	max := big.NewInt(int64(len(chars)))
	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return chars[num.Int64()], nil
}
