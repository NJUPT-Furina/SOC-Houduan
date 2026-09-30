package main

import (
	"fmt"
)

func main() {
	countName := "B26040721"
	password := "26040721"
	isBlocked := false
	weekLogincount := []int{5, 7, 4, 1, 0, 7, 4}
	if countName == "" {
		fmt.Println("用户名不能为空")
	}
	if len(password) < 6 {
		fmt.Println("密码长度不能小于6位")
	}
	if isBlocked {
		fmt.Println("账号风险拦截")
	}
	if len(weekLogincount) <= 0 {
		fmt.Println("登录次数格式错误")
	}

	fmt.Println("账号校验通过")
}
