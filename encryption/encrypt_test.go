package encryption

import (
	"os"
	"testing"
)

// TestGenerateAndSaveAESKey tests if the key generation and saving function works
func TestGenerateAndSaveAESKey(t *testing.T) {
	_, err := GenerateAndSaveAESKey()
	if err != nil {
		t.Errorf("Failed to generate and save AES key: %v", err)
	}
}

// TestEncryptFile tests if a file is encrypted properly
func TestEncryptFile(t *testing.T) {
	// Create a test file
	testFile := "test.txt"
	err := os.WriteFile(testFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	// Generate an AES key
	key, err := GenerateAndSaveAESKey()
	if err != nil {
		t.Fatalf("Failed to generate AES key: %v", err)
	}

	// Encrypt the file
	err = EncryptFile(testFile, key)
	if err != nil {
		t.Errorf("Failed to encrypt file: %v", err)
	}

	// Check if the encrypted file was created
	encFile := testFile + ".enc"
	if _, err := os.Stat(encFile); os.IsNotExist(err) {
		t.Errorf("Expected encrypted file %s, but it was not found", encFile)
	}
	defer os.Remove(encFile)
}

