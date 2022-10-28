/*
Copyright 2022 The ChengHongXi Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cipher

import (
	"bytes"
	"encoding/base64"

	"crypto/aes"
	"crypto/cipher"
)

// Config struct,Define AES KEY,IV
type Config struct {
	AesKey string
	AesIv  string
}

type NewCipher struct {
}

func New() *NewCipher {
	return &NewCipher{}
}

// Encrypt 加密
func (e *NewCipher) Encrypt(data []byte, config Config) (string, error) {
	block, err := aes.NewCipher([]byte(config.AesKey))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	originData := pad(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(config.AesIv))
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// Decrypt 解密
func (e *NewCipher) Decrypt(decryptText string, config Config) ([]byte, error) {
	decodeData, err := base64.StdEncoding.DecodeString(decryptText)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher([]byte(config.AesKey))
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(config.AesIv))
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
