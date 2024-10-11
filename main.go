package main

import (
  "crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func getUserInput() string {
	var inputFile string
	fmt.Print("Enter file path: ")
	fmt.Scanln(&inputFile) // Use Scanln and pass the address of the variable
	fmt.Println("Entered file path: ", inputFile)
	return inputFile
}

func generateAESKey() []byte {
	key := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(key)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Generated AES Key: %x\n", key)
	return key
}

func createEncryptedFileName(filename string) string {
	// Remove any existing extension and add ".lock"
	return filename + ".lock"
}

func encryptFile() error {
	key := generateAESKey()
	inputFile := getUserInput()

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Create a new GCM (Galois/Counter Mode) cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// Generate a nonce (number used once) for GCM encryption
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Step 1: Open the input file for reading
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("could not open input file: %v", err)
	}
	defer inFile.Close()

	// Read file contents
	fileContent, err := io.ReadAll(inFile)
	if err != nil {
		return fmt.Errorf("could not read input file: %v", err)
	}

	// Step 2: Encrypt the file content
	ciphertext := aesgcm.Seal(nonce, nonce, fileContent, nil)

	// Step 3: Create a new file with the ".enc" extension
	encFileName := createEncryptedFileName(inputFile)
	encFile, err := os.Create(encFileName) // Create the new encrypted file
	if err != nil {
		return fmt.Errorf("could not create encrypted file: %v", err)
	}
	defer encFile.Close()

	// Step 4: Write the encrypted content to the new file
	_, err = encFile.Write(ciphertext)
	if err != nil {
		return fmt.Errorf("could not write encrypted data to file: %v", err)
	}

	// Step 5: Close and remove the original plaintext file
	if err := os.Remove(inputFile); err != nil {
		return fmt.Errorf("could not delete the original file: %v", err)
	}

	fmt.Println("File encrypted successfully!")
	return nil
}

func main() {
	if err := encryptFile(); err != nil {
		fmt.Println("Error:", err)
	}
}

