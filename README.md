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
   - Saves the key securely for future decryption in `~/.keydir`.

2. **Decryption**:
   - Loads the previously saved AES key from `~/.keydir`.
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

## Running with Docker

You can also run the SecureVault tool using Docker. The container runs as a **non-root user**. This can be changed in the Dockerfile if needed.

### Pull from Docker Hub

1. **Pull the Docker image**:
   ```bash
   docker pull arviiyer/securevault:latest
   ```

2. **Run the Docker container**:
   ```bash
   docker run --rm -it -v /path/to/your/files:/data -v ~/.keydir:/home/secureuser/.keydir --user $(id -u):$(id -g) -e KEY_DIR="/home/secureuser/.keydir" arviiyer/securevault:latest
   ```

   - **Explanation**:
     - `-v /path/to/your/files:/data`: Mounts the directory containing the files to be encrypted or decrypted.
     - `-v ~/.keydir:/home/secureuser/.keydir`: Mounts the key directory to ensure the encryption key is accessible.
     - `--user $(id -u):$(id -g)`: Runs the container as the current user to avoid root ownership issues.
     - `-e KEY_DIR="/home/secureuser/.keydir"`: Sets the key directory environment variable inside the container.

---

## Example

### Encryption
```bash
Would you like to encrypt or decrypt? (e/d): e
Enter directory path: /data
Entered directory path:  /data
Encrypted /data/test.txt -> /data/test.txt.enc
Encrypted /data/test.txt successfully
All files in the directory were encrypted successfully!
```

### Decryption
```bash
Would you like to encrypt or decrypt? (e/d): d
Enter directory path: /data
Entered directory path:  /data
Decrypted /data/test.txt.enc -> /data/test.txt
Decrypted /data/test.txt.enc successfully
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
- **Parallel File Processing Enhancements**: Further optimize concurrency for even larger workloads.

---

## License

This project is licensed under the [MIT License](./LICENSE).

