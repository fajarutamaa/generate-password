package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/fajarutamaa/generate-password/password"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a secure password",
	RunE: func(cmd *cobra.Command, args []string) error {
		length, _ := cmd.Flags().GetInt("length")
		includeSymbols, _ := cmd.Flags().GetBool("symbols")
		includeNumbers, _ := cmd.Flags().GetBool("numbers")
		includeUppercase, _ := cmd.Flags().GetBool("uppercase")

		if length < 8 {
			return errors.New("length must be at least 8 characters")
		}

		pass, err := password.Generate(length, includeSymbols, includeNumbers, includeUppercase)
		if err != nil {
			return err
		}

		fmt.Println("Result generated password:", pass)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 12, "Length of the password")
	generateCmd.Flags().BoolP("symbols", "s", false, "Include symbols")
	generateCmd.Flags().BoolP("numbers", "n", false, "Include numbers")
	generateCmd.Flags().BoolP("uppercase", "u", false, "Include uppercase letters")
}
