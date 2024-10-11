package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// Function to generate AES-256 key
func GenerateAESKey() []byte {
	key := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(key)
	if err != nil {
		panic(err.Error())
	}
	return key
}

// Helper function to create the ".enc" file extension
func CreateEncryptedFileName(filename string) string {
	return filename + ".lock"
}

// Function to encrypt the file
func EncryptFile(inputFile string, key []byte) error {
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
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("could not open input file: %v", err)
	}
	defer inFile.Close()

	// Read the content of the file
	fileContent, err := io.ReadAll(inFile)
	if err != nil {
		return fmt.Errorf("could not read input file: %v", err)
	}

	// Encrypt the content
	ciphertext := aesgcm.Seal(nonce, nonce, fileContent, nil)

	// Create encrypted file
	encFileName := CreateEncryptedFileName(inputFile)
	outFile, err := os.Create(encFileName)
	if err != nil {
		return fmt.Errorf("could not create encrypted file: %v", err)
	}
	defer outFile.Close()

	// Write the encrypted content to the new file
	_, err = outFile.Write(ciphertext)
	if err != nil {
		return fmt.Errorf("could not write to encrypted file: %v", err)
	}

	// Remove the original file
	err = os.Remove(inputFile)
	if err != nil {
		return fmt.Errorf("could not remove original file: %v", err)
	}

	return nil
}

