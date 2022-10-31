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
