/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	symbolChars  = "!@#$%^&*()-_=+[]{}<>?/|"
	defaultChars = lowerChars + upperChars + numberChars
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a secure password",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		length, _ := cmd.Flags().GetInt("length")
		includeSymbols, _ := cmd.Flags().GetBool("symbols")
		includeNumbers, _ := cmd.Flags().GetBool("numbers")
		includeUppercase, _ := cmd.Flags().GetBool("uppercase")

		if length < 0 {
			return errors.New("length must be at least 8 characters")
		}

		password, err := generatePassword(length, includeSymbols, includeNumbers, includeUppercase)
		if err != nil {
			return err
		}

		cmd.Println("Result generated password:", password)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of the password")
	generateCmd.Flags().BoolP("symbols", "s", false, "Include symbols in the password")
	generateCmd.Flags().BoolP("numbers", "n", false, "Include numbers in the password")
	generateCmd.Flags().BoolP("uppercase", "u", false, "Include uppercase letters in the password")
}

func generatePassword(length int, includeSymbols bool, includeNumbers bool, includeUppercase bool) (string, error) {
	var password strings.Builder
	chars := lowerChars

	if includeSymbols {
		chars += upperChars
	}

	if includeNumbers {
		chars += numberChars
	}

	if includeUppercase {
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
