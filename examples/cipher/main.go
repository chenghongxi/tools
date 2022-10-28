package main

import (
	"fmt"

	"github.com/chenghongxi/tools/cipher"
)

type Ciphers struct {
	key string
	iv  string
}

func main() {

	// 自定义变量b 类型为：byte数组
	b := []byte("hello,world")

	// 自定义AES_KEY AES_IV
	AES_KEY := "KHGSI69YBWGS0TXB"
	AES_IV := "3010201735544646"

	// 加密b
	c := cipher.New()
	encrypt, err := c.Encrypt(b, AES_KEY, AES_IV)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("加密：", encrypt)

	// 解密
	decrypt, err := c.Decrypt(encrypt, AES_KEY, AES_IV)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密：", string(decrypt))

}
