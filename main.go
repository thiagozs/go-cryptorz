package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Cryptorz data
type Cryptorz struct {
	Key []byte
}

func main() {

	key := []byte("thiagozs-poc-of-concept;1234567@")

	cz := NewCryptorz(key)

	encoder := cz.ZEncrypt("Teste de encode/decoder")

	fmt.Printf("Encoder String: %s\n", encoder)

	decoder := cz.ZDecrypt(encoder)

	fmt.Printf("Decoder String: %s\n", decoder)

}

// NewCryptorz pass the key for enc and dec
func NewCryptorz(key []byte) Cryptorz {
	c := Cryptorz{}
	c.Key = key
	return c
}

//ZEncrypt string to base64 crypto using AES
func (c *Cryptorz) ZEncrypt(text string) string {
	plaintext := []byte(text)
	block, err := aes.NewCipher(c.Key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

//ZDecrypt from base64 to decrypted string
func (c *Cryptorz) ZDecrypt(crypttext string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(crypttext)
	block, err := aes.NewCipher(c.Key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	//can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
