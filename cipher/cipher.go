package cipher

import (
	"bytes"
	"encoding/base64"

	"crypto/aes"
	"crypto/cipher"
)

const (
	AES_KEY = "KHGSI69YBWGS0TXB"
	AES_IV  = "3010201735544646"
)

type Ciphers struct {
	AesKey string
	AesIv  string
}

type NewCipher struct {
}

func New() *NewCipher {
	return &NewCipher{}
}

func (e *NewCipher) Encrypt(data []byte, ciphers *Ciphers) (string, error) {
	block, err := aes.NewCipher([]byte(ciphers.AesKey))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	originData := pad(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(ciphers.AesIv))
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func (e *NewCipher) Decrypt(decryptText string, ciphers *Ciphers) ([]byte, error) {
	decodeData, err := base64.StdEncoding.DecodeString(decryptText)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher([]byte(ciphers.AesKey))
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(ciphers.AesIv))
	originData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(originData, decodeData)
	return unPad(originData), nil
}

func pad(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func unPad(cipherText []byte) []byte {
	length := len(cipherText)
	unPadding := int(cipherText[length-1])
	return cipherText[:(length - unPadding)]
}
