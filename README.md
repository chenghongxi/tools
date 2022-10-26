# tools
library and tool for golang

## config
- 完成 `yaml` 配置文件的加载
- 完成 `json` 配置文件的加载
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

