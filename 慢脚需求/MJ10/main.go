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
	if err := validateCount(c.status); err != nil {
		fmt.Println("验证账户失败:", err)
		return
	}
	if err := c.UpdatePassword("26040721cqx"); err != nil {
		fmt.Println("更新密码失败:", err)
		return
	}

	fmt.Println("更新密码成功")
}
