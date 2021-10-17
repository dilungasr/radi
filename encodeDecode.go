package radi

import (
	"encoding/base64"
)

//base64Encoder decodes the encrypted slice of byte to base64 string
func base64Encoder(phrase []byte) string {
	return base64.URLEncoding.EncodeToString(phrase)
}

// base64Decoder for decoding base64 string to the slice of bytes
func base64Decoder(phrase string) (data []byte, valid bool) {
	data, err := base64.URLEncoding.DecodeString(phrase)

	//if it is not a valid base64 string
	if err != nil {
		return data, false
	}

	return data, true
}
