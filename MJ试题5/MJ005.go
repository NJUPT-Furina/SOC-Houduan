package main

import (
	"fmt"
)

func main() {
	countName := []string{"cqx001", "ttdr001", "cqx002", "", "cqx003", "tangent_0712", "cqx114514", "risk_user", "cqx005", "cqx006"}
	count := 0
	for _, countName := range countName {
		if countName == "" {
			continue
		}
		if countName == "risk_user" {
			fmt.Println("账号风险拦截")
			break
		}

		count++
		fmt.Printf("正在创建账号: %s\n", countName)
	}
	fmt.Printf("总共创建了 %d 个账号\n", count)
}
