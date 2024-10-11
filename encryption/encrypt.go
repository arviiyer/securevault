package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GenerateAESKey generates a 256-bit AES key and saves it to a file
func GenerateAndSaveAESKey() ([]byte, error) {
	// Generate a random 256-bit AES key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate AES key: %v", err)
	}

	// Ensure the "key/" directory exists
	err = os.MkdirAll("key", 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create key directory: %v", err)
	}

	// Save the key to a file
	keyPath := filepath.Join("key", "aes.key")
	err = os.WriteFile(keyPath, key, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to save AES key: %v", err)
	}

	fmt.Println("AES key saved to:", keyPath)
	return key, nil
}

// EncryptFile encrypts a single file and saves it with a .enc extension
func EncryptFile(filePath string, key []byte) error {
	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create a new GCM (Galois/Counter Mode) cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Generate a nonce for encryption
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Open the input file for reading
	inFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file %v: %v", filePath, err)
	}
	defer inFile.Close()

	// Read the file content
	fileContent, err := io.ReadAll(inFile)
	if err != nil {
		return fmt.Errorf("could not read file %v: %v", filePath, err)
	}

	// Encrypt the content
	ciphertext := aesgcm.Seal(nonce, nonce, fileContent, nil)

	// Create the output file with the ".enc" extension
	encFilePath := filePath + ".enc"
	outFile, err := os.Create(encFilePath)
	if err != nil {
		return fmt.Errorf("could not create encrypted file %v: %v", encFilePath, err)
	}
	defer outFile.Close()

	// Write the encrypted content to the new file
	_, err = outFile.Write(ciphertext)
	if err != nil {
		return fmt.Errorf("could not write to encrypted file %v: %v", encFilePath, err)
	}

	// Remove the original file after encryption
	err = os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("could not delete original file %v: %v", filePath, err)
	}

	fmt.Printf("Encrypted %s -> %s\n", filePath, encFilePath)
	return nil
}

// EncryptFilesInDirectory encrypts all files in a given directory
func EncryptFilesInDirectory(dirPath string, key []byte) error {
	// Walk through all files in the directory
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a regular file (not a directory)
		if !info.IsDir() {
			// Encrypt the file
			err := EncryptFile(path, key)
			if err != nil {
				fmt.Printf("Failed to encrypt file %s: %v\n", path, err)
			}
		}
		return nil
	})
	return err
}

