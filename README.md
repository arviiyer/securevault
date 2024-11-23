# SecureVault: A Secure File Encryption and Decryption Tool

SecureVault is a lightweight tool built in Go that allows users to securely encrypt and decrypt files within a directory using AES-256 encryption. Now, with **concurrency**, SecureVault takes advantage of multi-core processing to make file encryption and decryption significantly faster. This tool is ideal for anyone looking to safeguard sensitive data or learn about implementing cryptographic techniques.

---

## ✨ Features

- 🔒 **AES-256 Encryption**: Industry-standard encryption to secure your files.
- 🚀 **Blazingly Fast**: Utilizes goroutines and worker pools to encrypt and decrypt multiple files concurrently.
- 📂 **Directory-Level Processing**: Encrypts or decrypts all files in a specified directory.
- 🔑 **Key Management**:
  - 🛠️ Automatically generates a secure encryption key.
  - 🔐 Saves and retrieves the key for seamless decryption.
- 🖥️ **Simple Command-Line Interface**: User-friendly prompts for encryption and decryption.

---

## How It Works

1. **Encryption**:
   - Generates a random AES-256 encryption key.
   - Encrypts all files in the specified directory concurrently and appends `.enc` to their filenames.
   - Saves the key securely for future decryption.

2. **Decryption**:
   - Loads the previously saved AES key.
   - Decrypts all `.enc` files in the specified directory concurrently and restores their original state.

---

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/arviiyer/securevault.git
   cd securevault
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build -o securevault
   ```

---

## Usage

1. **Run the program**:
   ```bash
   ./securevault
   ```

2. **Follow the prompts**:
   - Choose between encryption (`e`) or decryption (`d`).
   - Provide the path to the directory containing the files you wish to process.

---

## Example

### Encryption
```bash
Would you like to encrypt or decrypt? (e/d): e
Enter directory path: /path/to/your/files
All files in the directory were encrypted successfully!
```

### Decryption
```bash
Would you like to encrypt or decrypt? (e/d): d
Enter directory path: /path/to/your/files
All .enc files in the directory were decrypted successfully!
```

---

## Project Structure

- **main.go**: Entry point for the program, handling user input and directing actions.
- **encryption**:
  - `encrypt.go`: Handles key generation and file encryption.
- **decryption**:
  - `decrypt.go`: Manages key loading and file decryption.

---

## Requirements

- Go 1.20+ (or later)

---

## Future Enhancements

- Support for additional encryption algorithms.
- Integration with cloud storage for key backup.
- Logging and error reporting enhancements.
- Parallel File Processing Enhancements: Further optimize concurrency for even larger workloads.

---

## License

This project is licensed under the [MIT License](./LICENSE).

