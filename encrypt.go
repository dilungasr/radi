package radi

import (
	"crypto/aes"
	"crypto/cipher"
)

// Encrypt encrypts plainPhrase using the given secrete key and initialization vector(iv).
// Then returns an encrypted base64 URL-encoded string.
func Encrypt(plainPhrase, secreteKey string) (encryptedPhrase string, err error) {
	//  create aes cipher
	block, err := aes.NewCipher([]byte(secreteKey))
	if err != nil {
		return encryptedPhrase, err
	}

	// generate IV (initalization vector)
	IV, err := RandBytes(16)
	if err != nil {
		return encryptedPhrase, err
	}

	// perform encryption
	plainText := []byte(plainPhrase)
	cfb := cipher.NewCFBEncrypter(block, IV)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	// prepend IV to the ciphertext
	cipherText = append(IV, cipherText...)

	// return the base64Encoded string
	return base64Encoder(cipherText), err
}
