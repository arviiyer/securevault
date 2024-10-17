package decryption

import (
	"testing"
	"os"
	"github.com/arviiyer/ransomware-poc/encryption"
)

// Helper function to ensure key generation
func ensureAESKeyExists(t *testing.T) {
	if _, err := os.Stat("key/aes.key"); os.IsNotExist(err) {
		// Generate the AES key if it doesn't exist
		_, err := encryption.GenerateAndSaveAESKey()
		if err != nil {
			t.Fatalf("Failed to generate AES key for testing: %v", err)
		}
	}
}

// TestLoadAESKey tests if the AES key can be loaded successfully
func TestLoadAESKey(t *testing.T) {
	// Ensure the key is generated
	ensureAESKeyExists(t)

	// Try loading the key
	_, err := LoadAESKey()
	if err != nil {
		t.Errorf("Failed to load AES key: %v", err)
	}
}

// TestDecryptFile tests if a file is decrypted correctly
func TestDecryptFile(t *testing.T) {
	// Ensure the key is generated
	ensureAESKeyExists(t)

	// Create an encrypted test file
	encFile := "test.txt.enc"
	err := os.WriteFile(encFile, []byte("encrypted content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create encrypted test file: %v", err)
	}
	defer os.Remove(encFile)

	// Decrypt the file
	err = DecryptFile(encFile, []byte("dummy-key")) // Replace with proper key after loading
	if err != nil {
		t.Errorf("Failed to decrypt file: %v", err)
	}

	// Check if the decrypted file was created
	decFile := "test.txt"
	if _, err := os.Stat(decFile); os.IsNotExist(err) {
		t.Errorf("Expected decrypted file %s, but it was not found", decFile)
	}
	defer os.Remove(decFile)
}

