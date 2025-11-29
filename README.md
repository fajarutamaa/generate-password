# Generate Password CLI

A secure command-line password generator tool built with Go and Cobra CLI framework.

## 🌟 Features

- **Customizable Length**: Generate passwords with any length (minimum 8 characters)
- **Character Sets**: Include/exclude uppercase letters, numbers, and symbols
- **Password Strength Checker**: Automatic strength calculation with color-coded indicators
- **Multiple Password Generation**: Generate multiple passwords at once with `--count`
- **Clipboard Support**: Copy passwords directly to clipboard with `--copy`
- **Colored Output**: Visual feedback with emoji and ANSI colors
- **Shell Completion**: Auto-completion support for bash, zsh, fish, and powershell
- **Secure Generation**: Uses cryptographically secure random number generation (crypto/rand)
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Simple CLI**: Easy-to-use command-line interface with helpful error messages

## 📦 Installation

### Option 1: Install from source (recommended)

```bash
go install github.com/fajarutamaa/generate-password@latest
```

Make sure `$GOPATH/bin` is in your PATH:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Option 2: Build from source

```bash
# Clone the repository
git clone https://github.com/fajarutamaa/generate-password.git
cd generate-password

# Build and install
go build -o generate-password .
sudo mv generate-password /usr/local/bin/
```

## 🚀 Usage

### Basic Commands

```bash
# Show help
generate-password -h

# Show generate command help
generate-password generate -h
```

### Generate Passwords

```bash
# Generate default password (12 characters, lowercase only)
generate-password generate

# Generate password with custom length
generate-password generate --length 16
generate-password generate -l 20

# Generate password with numbers
generate-password generate --numbers
generate-password generate -n

# Generate password with uppercase letters
generate-password generate --uppercase
generate-password generate -u

# Generate password with symbols
generate-password generate --symbols
generate-password generate -s

# Generate complex password (all options)
generate-password generate -l 16 -s -n -u
generate-password generate --length 20 --symbols --numbers --uppercase

# Generate multiple passwords at once
generate-password generate --count 5 -l 16 -u -n -s
generate-password generate -c 3 -l 12 -u -n

# Copy password to clipboard
generate-password generate -l 16 -u -n -s --copy

# Generate without colors (for scripting)
generate-password generate --no-color
```

## 📝 Examples

```bash
# Simple 12-character password (lowercase only)
$ generate-password generate
🔴 Generated Password: abcdefghijkl
   Strength: Weak (Score: 25/100)

# 16-character password with all options (Very Strong)
$ generate-password generate -l 16 -u -n -s
🟦 Generated Password: KkZPW!Fat06#+Fz*
   Strength: Very Strong (Score: 85/100)

# Generate 5 passwords at once
$ generate-password generate --count 5 -l 12 -u -n

1. 🟡 Password: dVjEDSkXqdkp
   Strength: Medium (Score: 50/100)

2. 🟢 Password: AV2U4H46LeO2
   Strength: Strong (Score: 65/100)

3. 🟢 Password: D8wElWwTmsFI
   Strength: Strong (Score: 65/100)

4. 🟡 Password: hRmbColGCQRM
   Strength: Medium (Score: 50/100)

5. 🟢 Password: p9jEtIPp1vEw
   Strength: Strong (Score: 65/100)

# Perfect score password (24 characters, all options)
$ generate-password generate -l 24 -u -n -s
🟦 Generated Password: #!|7BnuiCwm)k@vn}(aC)f<6
   Strength: Very Strong (Score: 100/100)

# Copy to clipboard
$ generate-password generate -l 16 -u -n -s --copy
✓ Generated password copied to clipboard!
🟢 Generated Password: -ttMHJK%aDuvjGwH
   Strength: Strong (Score: 75/100)
```

## 🔧 Command Line Options

| Flag | Short | Description | Default |
|------|--------|-------------|---------|
| `--length` | `-l` | Length of the password (minimum: 8) | 12 |
| `--symbols` | `-s` | Include symbols (!@#$%^&*()-_=+[]{}<?>/\|) | false |
| `--numbers` | `-n` | Include numbers (0-9) | false |
| `--uppercase` | `-u` | Include uppercase letters (A-Z) | false |
| `--count` | `-c` | Number of passwords to generate | 1 |
| `--copy` | | Copy password to clipboard (only with --count=1) | false |
| `--no-color` | | Disable colored output | false |

## 🎨 Password Strength Indicators

Each generated password is automatically analyzed and assigned a strength level:

| Emoji | Level | Score Range | Description |
|-------|-------|-------------|-------------|
| 🔴 | Weak | 0-39 | Short length or limited character variety |
| 🟡 | Medium | 40-59 | Decent length with some character variety |
| 🟢 | Strong | 60-79 | Good length with multiple character types |
| 🟦 | Very Strong | 80-100 | Excellent length with all character types |

**Strength Calculation Factors:**
- **Length** (30%): Longer passwords score higher
- **Character Variety** (40%): Using uppercase, lowercase, numbers, and symbols
- **Entropy/Complexity** (30%): Charset size and overall complexity

## 🔐 Character Sets

- **Lowercase letters**: `abcdefghijklmnopqrstuvwxyz`
- **Uppercase letters**: `ABCDEFGHIJKLMNOPQRSTUVWXYZ` 
- **Numbers**: `0123456789`
- **Symbols**: `!@#$%^&*()-_=+[]{}<?>/|`

## 🛡️ Security

This tool uses Go's `crypto/rand` package which provides cryptographically secure random number generation, ensuring that generated passwords are truly random and secure.

## 🏗️ Project Structure

```
.
├── cmd/
│   ├── generate.go    # Generate command implementation
│   └── root.go        # Root command and CLI setup
├── password/
│   └── generate.go    # Password generation logic
├── main.go           # Main entry point
├── go.mod           # Go module definition
├── go.sum           # Go dependencies
├── LICENSE          # License file
└── README.md        # This file
```

## 🔨 Development

### Prerequisites

- Go 1.22.3 or higher
- Git

### Build from source

```bash
# Clone repository
git clone https://github.com/fajarutamaa/generate-password.git
cd generate-password

# Download dependencies
go mod tidy

# Build
go build -o generate-password .

# Test
./generate-password generate -h
```

### Run tests

```bash
go test ./...
```

## 🎯 Shell Completion

The CLI supports auto-completion for multiple shells. Generate and install completion scripts:

### Bash

```bash
# Generate completion script
generate-password completion bash > ~/.bash_completion.d/generate-password

# Or for system-wide installation
sudo generate-password completion bash > /etc/bash_completion.d/generate-password

# Then reload your shell
source ~/.bashrc
```

### Zsh

```bash
# Generate completion script
generate-password completion zsh > "${fpath[1]}/_generate-password"

# Then reload your shell
source ~/.zshrc
```

### Fish

```bash
# Generate completion script
generate-password completion fish > ~/.config/fish/completions/generate-password.fish

# Then reload your shell
source ~/.config/fish/config.fish
```

### PowerShell

```powershell
# Generate completion script
generate-password completion powershell | Out-String | Invoke-Expression

# To persist, add to your PowerShell profile
generate-password completion powershell >> $PROFILE
```

## 📄 Dependencies

- [spf13/cobra](https://github.com/spf13/cobra) - Modern CLI framework for Go
- [atotto/clipboard](https://github.com/atotto/clipboard) - Cross-platform clipboard library
- [fatih/color](https://github.com/fatih/color) - Colored terminal output

## 📝 License

This project is licensed under the [MIT License](LICENSE).

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📧 Contact

**Fajar Utama**
- GitHub: [@fajarutamaa](https://github.com/fajarutamaa)
- Email: [fajardwiutomo75@gmail.com](mailto:fajardwiutomo75@gmail.com)

## 🙏 Acknowledgments

- Thanks to the [Cobra CLI](https://cobra.dev/) team for the excellent framework
- Inspired by various password generation tools in the community