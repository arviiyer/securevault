package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// LoadAESKey loads the AES key from a file
func LoadAESKey() ([]byte, error) {
	keyFile, err := os.Open("/keys/aes_key")
	if err != nil {
		return nil, fmt.Errorf("could not open key file: %v", err)
	}
	defer keyFile.Close()

	encodedKey, err := io.ReadAll(keyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read key from file: %v", err)
	}

	key, err := base64.StdEncoding.DecodeString(strings.TrimSpace(string(encodedKey)))
	if err != nil {
		return nil, fmt.Errorf("could not decode key: %v", err)
	}

	return key, nil
}

// DecryptFile decrypts the given encrypted file with the provided AES key
func DecryptFile(filePath string, key []byte) error {
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

	// Open the input encrypted file for reading
	inFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open encrypted file %v: %v", filePath, err)
	}
	defer inFile.Close()

	// Read the encrypted content
	ciphertext, err := io.ReadAll(inFile)
	if err != nil {
		return fmt.Errorf("could not read encrypted file %v: %v", filePath, err)
	}

	// Extract the nonce from the beginning of the ciphertext
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return fmt.Errorf("ciphertext too short for file %v", filePath)
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the content
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("could not decrypt file %v: %v", filePath, err)
	}

	// Create the output file by removing the ".enc" extension
	decFilePath := strings.TrimSuffix(filePath, ".enc")
	outFile, err := os.Create(decFilePath)
	if err != nil {
		return fmt.Errorf("could not create decrypted file %v: %v", decFilePath, err)
	}
	defer outFile.Close()

	// Write the decrypted content to the new file
	_, err = outFile.Write(plaintext)
	if err != nil {
		return fmt.Errorf("could not write to decrypted file %v: %v", decFilePath, err)
	}

	// Remove the encrypted file after decryption
	err = os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("could not delete encrypted file %v: %v", filePath, err)
	}

	fmt.Printf("Decrypted %s -> %s\n", filePath, decFilePath)
	return nil
}

// DecryptFilesInDirectory decrypts all .enc files in a given directory
func DecryptFilesInDirectory(dirPath string, key []byte) error {
	// Walk through all files in the directory
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a regular file with the ".enc" extension
		if !info.IsDir() && strings.HasSuffix(path, ".enc") {
			// Decrypt the file
			err := DecryptFile(path, key)
			if err != nil {
				fmt.Printf("Failed to decrypt file %s: %v\n", path, err)
			}
		}
		return nil
	})
	return err
}
