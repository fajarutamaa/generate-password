# Generate Password CLI

A secure command-line password generator tool built with Go and Cobra CLI framework.

## 🌟 Features

- **Customizable Length**: Generate passwords with any length (minimum 8 characters)
- **Character Sets**: Include/exclude uppercase letters, numbers, and symbols
- **Secure Generation**: Uses cryptographically secure random number generation
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Simple CLI**: Easy-to-use command-line interface

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
```

## 📝 Examples

```bash
# Simple 12-character password
$ generate-password generate
Result generated password: abcdefghijkl

# 16-character password with numbers and symbols
$ generate-password generate -l 16 -n -s
Result generated password: a3b!c2d@e1f#g4h$

# Strong 20-character password with all options
$ generate-password generate --length 20 --symbols --numbers --uppercase
Result generated password: A3b!C2d@E1f#G4h$I5j%
```

## 🔧 Command Line Options

| Flag | Short | Description | Default |
|------|--------|-------------|---------|
| `--length` | `-l` | Length of the password (min: 8) | 12 |
| `--symbols` | `-s` | Include symbols (!@#$%^&*()-_=+[]{}<>?/\|) | false |
| `--numbers` | `-n` | Include numbers (0-9) | false |
| `--uppercase` | `-u` | Include uppercase letters (A-Z) | false |

## 🔐 Character Sets

- **Lowercase letters**: `abcdefghijklmnopqrstuvwxyz`
- **Uppercase letters**: `ABCDEFGHIJKLMNOPQRSTUVWXYZ` 
- **Numbers**: `0123456789`
- **Symbols**: `!@#$%^&*()-_=+[]{}<>?/|`

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

## 📄 Dependencies

- [spf13/cobra](https://github.com/spf13/cobra) - Modern CLI framework for Go

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