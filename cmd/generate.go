package cmd

import (
	"errors"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/fajarutamaa/generate-password/password"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a secure password",
	Long: `Generate one or more secure random passwords with customizable options.
	
You can control the length, character types, and generate multiple passwords at once.
The strength of each password is automatically calculated and displayed.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		length, _ := cmd.Flags().GetInt("length")
		includeSymbols, _ := cmd.Flags().GetBool("symbols")
		includeNumbers, _ := cmd.Flags().GetBool("numbers")
		includeUppercase, _ := cmd.Flags().GetBool("uppercase")
		count, _ := cmd.Flags().GetInt("count")
		copyToClipboard, _ := cmd.Flags().GetBool("copy")
		noColor, _ := cmd.Flags().GetBool("no-color")

		// Validation
		if length < 8 {
			return errors.New("length must be at least 8 characters")
		}

		if count < 1 {
			return errors.New("count must be at least 1")
		}

		if copyToClipboard && count > 1 {
			return errors.New("the --copy flag can only be used when generating a single password (--count=1)")
		}

		// Disable color if requested
		color.NoColor = noColor

		// Setup colors for strength levels
		red := color.New(color.FgRed).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()
		bold := color.New(color.Bold).SprintFunc()

		// Generate passwords
		for i := 0; i < count; i++ {
			pass, err := password.Generate(length, includeSymbols, includeNumbers, includeUppercase)
			if err != nil {
				return err
			}

			// Calculate strength
			score, level := password.CalculateStrength(pass, length)

			// Determine color based on strength
			var strengthColor func(a ...interface{}) string
			var emoji string
			switch level {
			case "Weak":
				strengthColor = red
				emoji = "🔴"
			case "Medium":
				strengthColor = yellow
				emoji = "🟡"
			case "Strong":
				strengthColor = green
				emoji = "🟢"
			case "Very Strong":
				strengthColor = cyan
				emoji = "🟦"
			default:
				strengthColor = green
				emoji = "🟢"
			}

			// Copy to clipboard if requested
			if copyToClipboard {
				err := clipboard.WriteAll(pass)
				if err != nil {
					return fmt.Errorf("failed to copy to clipboard: %w", err)
				}
				fmt.Println(bold("✓ Generated password copied to clipboard!"))
			}

			// Output format
			if count > 1 {
				fmt.Printf("\n%d. %s %s: %s\n", i+1, emoji, bold("Password"), pass)
			} else {
				fmt.Printf("%s %s: %s\n", emoji, bold("Generated Password"), pass)
			}

			fmt.Printf("   %s: %s (%s: %d/100)\n",
				bold("Strength"),
				strengthColor(level),
				bold("Score"),
				score,
			)
		}

		if count > 1 {
			fmt.Println() // Extra newline for better readability
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 12, "Length of the password (minimum: 8)")
	generateCmd.Flags().BoolP("symbols", "s", false, "Include symbols (!@#$%^&*()-_=+[]{}<?>/|)")
	generateCmd.Flags().BoolP("numbers", "n", false, "Include numbers (0-9)")
	generateCmd.Flags().BoolP("uppercase", "u", false, "Include uppercase letters (A-Z)")
	generateCmd.Flags().IntP("count", "c", 1, "Number of passwords to generate")
	generateCmd.Flags().Bool("copy", false, "Copy the generated password to clipboard (only works with --count=1)")
	generateCmd.Flags().Bool("no-color", false, "Disable colored output")
}
