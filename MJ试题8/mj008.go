package main

import (
	"fmt"
)

type Count struct {
	CountName string
	Name      string
	Password  string
	IsBlocked bool
	Number    int
}

func main() {
	count := Count{
		CountName: "B26040721",
		Name:      "cqx",
		Password:  "26040721cqx",
		IsBlocked: false,
		Number:    001,
	}
	count.Password = encryptPassword(count.Password)
	validateCountStruct(count.CountName, count.Password)
	fmt.Printf("账号信息: %+v\n", count)
}

func encryptPassword(password string) string {
	runes := []rune(password)
	for i := 0; i < len(runes); i++ {
		runes[i] = runes[i] + 1
	}
	return string(runes)
}

func validateCountStruct(countName string, password string) {
	if countName == "" && len(password) < 6 {
		fmt.Println("账号名或密码不合法！")
	}
}

func validateCount(isBlocked bool) {
	if isBlocked {
		fmt.Println("账号风险拦截")
	}
}
