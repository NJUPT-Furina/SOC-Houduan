package main

import (
	"errors"
	"fmt"
)

func main() {
	countName := "B26040721"
	password := "26040721"
	if err := validateCount(countName, password); err != nil {
		fmt.Println("验证账户失败:", err)
		return
	}
	password = encryptPassword(password)
	buildWelcomeMessage(countName)
}
func validateCount(countName string, password string) error {
	if countName == "" || len(password) < 6 {
		return errors.New("账号名或密码不合法！")
	}
	return nil
}
func encryptPassword(password string) string {
	runes := []rune(password)
	for i := 0; i < len(runes); i++ {
		runes[i] = runes[i] + 1
	}
	return string(runes)
}
func buildWelcomeMessage(countName string) {
	fmt.Printf("Hello,world!账号:%s诞生了\n", countName)
}
