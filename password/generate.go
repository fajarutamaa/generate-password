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

// CalculateStrength calculates the strength of a password
// Returns a score (0-100) and a level (Weak, Medium, Strong, Very Strong)
func CalculateStrength(password string, length int) (int, string) {
	score := 0

	// Length score (30 points max)
	if length >= 8 {
		score += 10
	}
	if length >= 12 {
		score += 10
	}
	if length >= 16 {
		score += 5
	}
	if length >= 20 {
		score += 5
	}

	// Character variety (40 points max - 10 each)
	hasLower := false
	hasUpper := false
	hasNumber := false
	hasSymbol := false

	for _, char := range password {
		switch {
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case strings.ContainsRune(symbolChars, char):
			hasSymbol = true
		}
	}

	if hasLower {
		score += 10
	}
	if hasUpper {
		score += 10
	}
	if hasNumber {
		score += 10
	}
	if hasSymbol {
		score += 10
	}

	// Entropy/complexity bonus (30 points max)
	charsetSize := 0
	if hasLower {
		charsetSize += 26
	}
	if hasUpper {
		charsetSize += 26
	}
	if hasNumber {
		charsetSize += 10
	}
	if hasSymbol {
		charsetSize += len(symbolChars)
	}

	// Calculate entropy bonus
	if charsetSize > 0 {
		// Entropy = length * log2(charsetSize)
		// We'll give points based on charset diversity and length
		if charsetSize >= 62 && length >= 12 { // all types except symbols
			score += 15
		} else if charsetSize >= 36 && length >= 10 { // at least 2 types
			score += 10
		} else if charsetSize >= 26 { // just one type
			score += 5
		}

		// Additional bonus for very long passwords
		if length >= 24 {
			score += 15
		} else if length >= 20 {
			score += 10
		} else if length >= 16 {
			score += 5
		}
	}

	// Cap at 100
	if score > 100 {
		score = 100
	}

	// Determine level
	var level string
	switch {
	case score < 40:
		level = "Weak"
	case score < 60:
		level = "Medium"
	case score < 80:
		level = "Strong"
	default:
		level = "Very Strong"
	}

	return score, level
}

func getRandom(chars string) (byte, error) {
	maxLen := big.NewInt(int64(len(chars)))
	num, err := rand.Int(rand.Reader, maxLen)
	if err != nil {
		return 0, err
	}
	return chars[num.Int64()], nil
}
