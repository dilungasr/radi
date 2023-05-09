package radi

import (
	"crypto/rand"
	"encoding/base64"
)

// RandString generates a random base64 URL-encoded string of length len
func RandString(len int) (randomString string, err error) {
	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return randomString, err
	}

	// return base64 Encodeed string of the length provided by the caller
	return base64.StdEncoding.EncodeToString(randomBytes)[:len], err
}

// RandBytes generates a random []byte of length len
//
// (Very handy for creating initiliazation vectors but you are not limited to!)
func RandBytes(len int) (randomBytes []byte, err error) {
	randBytes := make([]byte, 32)

	_, err = rand.Read(randBytes)
	if err != nil {
		return randBytes, err
	}

	return randBytes[:len], err
}

// RandDigits generates secure random digits(0-9) which can be used for anything
// from generating OTP to verification codes
func RandDigits(len int) (string, error) {
	codes := make([]byte, len)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < len; i++ {
		codes[i] = 48 + (codes[i] % 10)
	}

	return string(codes), nil
}
