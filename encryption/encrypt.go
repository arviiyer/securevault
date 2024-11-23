package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GenerateAndSaveAESKey generates a random AES-256 key and saves it to a file
func GenerateAndSaveAESKey() ([]byte, error) {
	key := make([]byte, 32) // 256 bits
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("could not generate key: %v", err)
	}

	encodedKey := base64.StdEncoding.EncodeToString(key)

	// Set the key directory to ~/.securevault/keydir
	keyDir := filepath.Join(os.Getenv("HOME"), ".securevault", "keydir")

	// Create key directory if it does not exist
	if err := os.MkdirAll(keyDir, 0600); err != nil {
		return nil, fmt.Errorf("could not create key directory: %v", err)
	}

	keyFilePath := filepath.Join(keyDir, "aes_key")
	keyFile, err := os.Create(keyFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not create key file: %v", err)
	}
	defer keyFile.Close()

	_, err = keyFile.WriteString(encodedKey)
	if err != nil {
		return nil, fmt.Errorf("could not write key to file: %v", err)
	}

	return key, nil
}

// EncryptFile encrypts the given file with the provided AES key
func EncryptFile(filePath string, key []byte) error {
	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create GCM (Galois/Counter Mode) from the AES block cipher
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
