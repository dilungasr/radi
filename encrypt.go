package radi

import (
	"crypto/aes"
	"crypto/cipher"
)

// Encrypt encrypts plainPhrase using the given secrete key and initialization vector(iv).
//Then returns an encrypted base64 URL-encoded string.
func Encrypt(plainPhrase, secreteKey string, iv []byte) (encryptedPhrase string) {
	//  create aes cipher
	block, err := aes.NewCipher([]byte(secreteKey))
	if err != nil {
		panic(err)
	}

	// convert the phrase to []bytes
	plainText := []byte(plainPhrase)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	// return the base64Encoded string
	return base64Encoder(cipherText)
}
