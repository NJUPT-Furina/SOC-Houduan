package main

import (
	"fmt"
	"os"
)

type Count struct {
	CountName    string
	Name         string
	Password     string
	IsnotBlocked bool
	Number       int
}

func main() {
	count := Count{
		CountName:    "B26040721",
		Name:         "cqx",
		Password:     "26040721cqx",
		IsnotBlocked: true,
		Number:       001,
	}
	count.Password = encryptPassword(count.Password)
	validateCountStruct(count.CountName, count.Password)
	fmt.Printf("账号名: %s\n姓名: %s\n密码: %s\n帐号状态: %t\n账户序号：%03d\n", count.CountName, count.Name, count.Password, count.IsnotBlocked, count.Number)
}

func encryptPassword(password string) string {
	runes := []rune(password)
	for i := 0; i < len(runes); i++ {
		runes[i] = runes[i] + 1
	}
	return string(runes)
}

func validateCountStruct(countName string, password string) {
	if countName == "" || len(password) < 6 {
		fmt.Println("账号名或密码不合法！")
		os.Exit(1)
	}
}

func validateCount(isnotBlocked bool) {
	if !isnotBlocked {
		fmt.Println("账号风险拦截")
		os.Exit(1)
	}

}
