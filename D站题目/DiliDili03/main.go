package main

import (
	"fmt"
)

func main() {
	users := []string{"普通用户", "大会员用户"}
	for _, user := range users {
		if user == "大会员用户" {
			fmt.Println("观看VIP视频")
		}

		fmt.Println("观看普通视频")
	}
}
