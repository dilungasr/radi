package radi

import (
	"crypto/aes"
	"crypto/cipher"
)

// Decrypt extracts the text from the base64 encoded encrypted string using the secrete key
// and initialization vectory used to encrypt the text
//
//plainPhrase - is a plain text extracted from an encrypted text
//
//valid - true if plainPhrase is a valid base64 string. It's false if plainPhrase is not a valid base64 string
func Decrypter(encryptedPhrase, secreteKey string, iv []byte) (plainPhrase string, valid bool) {
	// create the aes cipher block
	block, err := aes.NewCipher([]byte(secreteKey))
	if err != nil {
		panic(err)
	}

	// decode the given base64 string to slice of bytes
	cipherText, valid := base64Decoder(encryptedPhrase)

	if !valid {
		return "Invalid base64 string", false
	}

	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	// return the decrypted
	return string(plainText), true
}
