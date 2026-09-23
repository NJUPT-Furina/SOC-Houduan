package main

import (
	"fmt"
)

func main() {
	c, _ := NewCount("cqx001", "26040721cqx")
	validateCount(c.status)
	err := c.UpdatePassword("cqx26040721")
	if err != nil {
		fmt.Println("更新密码失败:", err)
	}
	fmt.Println("更新密码成功")
}
