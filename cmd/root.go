/*
Copyright © 2024 Fajar Utama <fajardwiutomo75@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "generate-password",
	Short: "A secure command-line password generator",
	Long: `Generate Password CLI - A secure password generator tool built with Go.

This tool helps you create strong, random passwords with customizable options
including length, character types, and more. Each password's strength is
automatically calculated and displayed.

Features:
  • Generate single or multiple passwords at once
  • Customize length and character types (uppercase, numbers, symbols)
  • Password strength checker with color-coded indicators
  • Copy passwords directly to clipboard
  • Shell completion support for bash, zsh, fish, and powershell

Examples:
  # Generate a simple password
  generate-password generate

  # Generate a strong 20-character password with all character types
  generate-password generate -l 20 -u -n -s

  # Generate 5 passwords at once
  generate-password generate --count 5 -l 16 -u -n -s

  # Generate and copy to clipboard
  generate-password generate -l 16 -u -n -s --copy

  # Generate shell completion
  generate-password completion bash > ~/.bash_completion.d/generate-password`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra built-in completion command is automatically available
	// No additional flags needed
}
