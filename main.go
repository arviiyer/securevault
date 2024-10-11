package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func getUserInput() string {
  var input_file string
  fmt.Print("\nEnter file path: ")
  fmt.Scan(&input_file)
  fmt.Println("Entered file path: ", input_file)
  return input_file
}

func generateAESKey() ([]byte){
  key := make([]byte, 32) //32 bytes = 256 bits
  _, err := rand.Read(key)
  if err!= nil {
    panic(err.Error())
  }
  fmt.Printf("Generated AES Key: %x", key)
  return key
}

func encryptFile() error {
  key := generateAESKey()
  input_file := getUserInput()
  
  // Create AES cipher block
  block, err := aes.NewCipher(key)
  if err != nil {
    panic(err.Error())
  }

  // Create a new GCM (Galois/Counter Mode) cipher
  aesgcm, err := cipher.NewGCM(block)
  if err != nil {
    panic(err.Error())
  }

  // Generate a nonce (number used once) for GCM encryption
  nonce := make([]byte, aesgcm.NonceSize())
  if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
    panic(err.Error())
  }

  // Open the input file
  inFile, err := os.Open(input_file)
  if err != nil {
    panic(err.Error())
  }
  defer inFile.Close()

  // Read file contents
  file_content, err := io.ReadAll(inFile)
  if err != nil {
    panic(err.Error())
  }

  //Encrypt the file content
  ciphertext := aesgcm.Seal(nonce, nonce, file_content, nil)

  // Reopen the same file for writing (overwrite mode)
  outFile, err := os.Create(input_file)
  if err != nil {
    panic(err.Error())
  }
  defer outFile.Close()
  
  // Write the encrypted content back to the input file
  _, err = outFile.Write(ciphertext)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println("File encrypted successfully!")
  return nil
}

func main() {
  if err := encryptFile(); err != nil {
    panic(err.Error())
  }
}
