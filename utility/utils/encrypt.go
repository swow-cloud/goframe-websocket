package utils

import (
	"encoding/hex"

	"github.com/gogf/gf/v2/crypto/gaes"
)

func Encrypt(string []byte, key []byte) (str string, err error) {
	encrypt, err := gaes.Encrypt(string, key)
	if err != nil {
		return "", err
	}
	str = hex.EncodeToString(encrypt)
	return
}

func Decrypt(plainText string, key []byte) (str string, err error) {
	decryptStr, _ := hex.DecodeString(plainText)
	strByte, err := gaes.Decrypt(decryptStr, key)
	if err != nil {
		return "", err
	}
	str = string(strByte)
	return
}
