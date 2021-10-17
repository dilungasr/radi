package radi

import (
	"crypto/rand"
	"encoding/base64"
)

//RandString generates a random base64 URL-encoded string of length len
func RandString(len int) (randomString string) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// return base64 Encodeed string of the length provided by the caller
	return base64.StdEncoding.EncodeToString(randomBytes)[:len]
}

// RandBytes generates a random []byte of length len
//
// (Very handy for creating initiliazation vectors but you are not limited to!)
func RandBytes(len int) (randomBytes []byte) {
	randBytes := make([]byte, 32)

	_, err := rand.Read(randBytes)
	if err != nil {
		panic(err)
	}

	return randBytes[:len]
}
