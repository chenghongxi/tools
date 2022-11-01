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
