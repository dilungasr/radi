package radi

import (
	"encoding/base64"
)

// base64Encoder decodes the encrypted slice of byte to base64 string
func base64Encoder(phrase []byte) string {
	return base64.URLEncoding.EncodeToString(phrase)
}

// base64Decoder for decoding base64 string to the slice of bytes
func base64Decoder(phrase string) (data []byte, err error) {
	data, err = base64.URLEncoding.DecodeString(phrase)

	return data, err
}
