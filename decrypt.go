package radi

import (
	"crypto/aes"
	"crypto/cipher"
)

// Decrypt extracts the text from the base64 encoded encrypted string using the secrete key and initialization vector used to encrypt the text
//
// plainPhrase - is a plain text extracted from an encrypted text
//
// valid - true if plainPhrase is a valid base64 string. It's false if plainPhrase is not a valid base64 string
func Decrypt(encryptedPhrase, secreteKey string) (plainPhrase string, err error) {
	// create the aes cipher block
	block, err := aes.NewCipher([]byte(secreteKey))
	if err != nil {
		return plainPhrase, err
	}

	// decode the given base64 string to slice of bytes
	cipherText, err := base64Decoder(encryptedPhrase)
	if err != nil {
		return plainPhrase, err
	}

	// separate IV from cipherText
	IV := cipherText[:16]
	cipherText = cipherText[16:]

	// perform decryption
	cfb := cipher.NewCFBDecrypter(block, IV)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	// return the decrypted text
	return string(plainText), err
}
