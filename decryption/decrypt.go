package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// LoadAESKey loads the AES key from the file in the "key/" directory
func LoadAESKey() ([]byte, error) {
	keyPath := filepath.Join("key", "aes.key")
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load AES key: %v", err)
	}
	return key, nil
}

// DecryptFile decrypts a single .enc file and restores it to its original form
func DecryptFile(encFilePath string, key []byte) error {
	// Open the encrypted file for reading
	inFile, err := os.Open(encFilePath)
	if err != nil {
		return fmt.Errorf("could not open encrypted file %v: %v", encFilePath, err)
	}
	defer inFile.Close()

	// Read the nonce (the first part of the file)
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("could not create cipher block: %v", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("could not create GCM: %v", err)
	}

	nonceSize := aesgcm.NonceSize()
	nonce := make([]byte, nonceSize)

	_, err = io.ReadFull(inFile, nonce)
	if err != nil {
		return fmt.Errorf("could not read nonce from file: %v", err)
	}

	// Read the rest of the file (the ciphertext and tag)
	ciphertext, err := io.ReadAll(inFile)
	if err != nil {
		return fmt.Errorf("could not read ciphertext from file: %v", err)
	}

	// Decrypt the file content
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("decryption failed: %v", err)
	}

	// Create the output file by removing the .enc extension
	origFilePath := strings.TrimSuffix(encFilePath, ".enc")
	outFile, err := os.Create(origFilePath)
	if err != nil {
		return fmt.Errorf("could not create decrypted file %v: %v", origFilePath, err)
	}
	defer outFile.Close()

	// Write the decrypted content to the output file
	_, err = outFile.Write(plaintext)
	if err != nil {
		return fmt.Errorf("could not write to decrypted file %v: %v", origFilePath, err)
	}

	// Remove the encrypted file after successful decryption
	err = os.Remove(encFilePath)
	if err != nil {
		return fmt.Errorf("could not delete encrypted file %v: %v", encFilePath, err)
	}

	fmt.Printf("Decrypted %s -> %s\n", encFilePath, origFilePath)
	return nil
}

// DecryptFilesInDirectory decrypts all .enc files in a given directory
func DecryptFilesInDirectory(dirPath string, key []byte) error {
	// Walk through all files in the directory
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only decrypt files with the .enc extension
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".enc") {
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

