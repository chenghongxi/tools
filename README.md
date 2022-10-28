# tools
library and tools for golang

## config
- 完成 `yaml` 配置文件的加载
- 完成 `json` 配置文件的加载

### 如何使用:
```go
package main

import (
	"fmt"

	"github.com/chenghongxi/tools/config"
)

type YamlConfig struct {
	Default YamlDefaultOption `yaml:"default"`
}

type YamlDefaultOption struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

type JsonConfig struct {
	Default JsonDefaultOption `json:"default"`
}

type JsonDefaultOption struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	// 解析 yaml文件
	c := config.New()
	c.SetConfigFile("yaml/config.yaml")
	c.SetConfigType("yaml")

	var yaml YamlConfig
	if err := c.Binding(&yaml); err != nil {
		panic(err)
	}
	fmt.Println(yaml.Default.Age)

	// 解析 json文件
	d := config.New()
	d.SetConfigFile("json/config.json")
	c.SetConfigType("json")
	var json JsonConfig
	if err := c.Binding(&json); err != nil {
		panic(err)
	}
	fmt.Println(json.Default.Age)

}
```
[Code_Demo](./examples/config/main.go)

## cipher
- 完成 加密与解密

### 如何使用:
```go
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
```
### output:
```shell
LFUbTDASAOmpv0iA1org8g==
hello,world
```
[Code_Demo](./examples/cipher/main.go)
