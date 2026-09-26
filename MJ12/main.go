package main

import (
	"fmt"
	"strings"
)

func main() {
	countName := " cqx001 "
	password := "26040721cqx"
	cleanCountName := strings.TrimSpace(countName)

	if isEmptyCount(countName) {
		fmt.Println("账号名不能为空！")
		return
	}
	if cleanCountName == password {
		fmt.Println("账号名不能与密码相同！")
		return
	}
	if strings.Contains(cleanCountName, "退款") {
		safeCleanCountName := strings.ReplaceAll(cleanCountName, "退款", "**")
		fmt.Println("Hello world!欢迎：", safeCleanCountName)
		return
	}
	fmt.Println("Hello world!欢迎：", cleanCountName)
}

func isEmptyCount(countName string) bool {
	return strings.TrimSpace(countName) == ""
}
