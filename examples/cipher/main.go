package main

import (
	"fmt"

	"github.com/chenghongxi/tools/cipher"
)

func main() {

	// 自定义变量b 类型为：byte数组
	b := []byte("hello,world")

	NewConfig := cipher.Config{
		AES_KEY: "KHGSI69YBWGS0TXB",
		AES_IV:  "3010201735544646",
	}

	// 加密b
	c := cipher.New()
	encrypt, err := c.Encrypt(b, NewConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("加密：", encrypt)

	// 解密
	decrypt, err := c.Decrypt(encrypt, NewConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密：", string(decrypt))

}
