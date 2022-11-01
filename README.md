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

type IniConfig struct {
	Default IniDefaultOption `ini:"default"`
}

type IniDefaultOption struct {
	Name string `ini:"name"`
	Age  int    `ini:"age"`
}

func main() {

	// 解析 yaml文件
	c := config.New()
	c.SetConfigFile("examples/config/yaml/config.yaml")
	c.SetConfigType("yaml")

	var yaml YamlConfig
	if err := c.Binding(&yaml); err != nil {
		panic(err)
	}
	fmt.Println(yaml.Default.Age)

	// 解析 json文件
	d := config.New()
	d.SetConfigFile("examples/config/json/config.json")
	c.SetConfigType("json")
	var json JsonConfig
	if err := c.Binding(&json); err != nil {
		panic(err)
	}
	fmt.Println(json.Default.Age)

	// 解析 ini文件
	i := config.New()
	i.SetConfigFile("examples/config/ini/my.ini")
	i.SetConfigType("ini")
	var ini IniConfig
	if err := c.Binding(&ini); err != nil {
		panic(err)
	}
	fmt.Println(ini.Default.Age)
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

func main() {

	// 自定义变量b 类型为：byte数组
	b := []byte("hello,world")

	NewConfig := cipher.Config{
		AesKey: "KHGSI69YBWGS0TXB",
		AesIv:  "3010201735544646",
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
```
### output:
```shell
LFUbTDASAOmpv0iA1org8g==
hello,world
```
[Code_Demo](./examples/cipher/main.go)

## exec
### 如何使用:
```go
package main

import (
	"fmt"

	"github.com/caoyingjunz/pixiulib/exec"
)

func main() {
	exec := exec.New()

	// 确认命令行是否存在
	if _, err := exec.LookPath("ping"); err != nil {
		panic(err)
	}
	// 属性
	out, err := exec.Command("ping", "www.baidu.com").CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}

```
### output:
```shell
���� Ping www.a.shifen.com [39.156.66.18] ���� 32 �ֽڵ�����:
���� 39.156.66.18 �Ļظ�: �ֽ�=32 ʱ��=27ms TTL=51
���� 39.156.66.18 �Ļظ�: �ֽ�=32 ʱ��=24ms TTL=51
���� 39.156.66.18 �Ļظ�: �ֽ�=32 ʱ��=24ms TTL=51
���� 39.156.66.18 �Ļظ�: �ֽ�=32 ʱ��=29ms TTL=51

39.156.66.18 �� Ping ͳ����Ϣ:
    ���ݰ�: �ѷ��� = 4���ѽ��� = 4����ʧ = 0 (0% ��ʧ)��
�����г̵Ĺ���ʱ��(�Ժ���Ϊ��λ):
    ���� = 24ms��� = 29ms��ƽ�� = 26ms


```

