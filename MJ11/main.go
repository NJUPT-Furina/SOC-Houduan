package main

import (
	"fmt"
)

func main() {
	c, err := NewCount("cqx001", "26040721cqx")
	if err != nil {
		fmt.Println("创建账户失败:", err)
		return
	}
	fmt.Println("Hello world! 欢迎：", c.CountName)
}
