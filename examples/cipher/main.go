package main

import (
	"fmt"

	"github.com/chenghongxi/tools/cipher"
)

func main() {

	// 定义变量b 类型为：byte数组
	b := []byte("hello,world")

	// 加密b
	c := cipher.New()
	encrypt, err := c.Encrypt(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("加密：", encrypt)

	// 解密
	decrypt, err := c.Decrypt(encrypt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密：", string(decrypt))

}
