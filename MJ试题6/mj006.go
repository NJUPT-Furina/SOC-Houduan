package main

import (
	"fmt"
)

func main() {
	countName := "B26040721"
	password := "26040721cqx"
	validateCount(countName, password)
	password = encryptPassword(password)
	buildWelcomeMessage(countName)
}
func validateCount(countName string, password string) {
	if countName == "" && len(password) < 6 {
		fmt.Println("账号名或密码不合法！")
	}
}
func encryptPassword(password string) string {
	runes := []rune(password)
	for i := 0; i < len(runes); i++ {
		runes[i] = runes[i] + 1
	}
	return string(runes)
}
func buildWelcomeMessage(countName string) {
	fmt.Printf("Hello,world!慢脚账号创建完成，账号:%s诞生了\n", countName)
}
