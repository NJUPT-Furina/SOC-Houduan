package main

import (
	"fmt"
	"os"
)

func main() {
	c, err := NewCount("cqx001", "弄臣", "26040721cqx")
	if err != nil {
		fmt.Println("创建账户失败:", err)
		return
	}
	file, err := os.Create("slowfoot_accounts.txt")
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(c.String())
	if err != nil {
		fmt.Println("写入失败：", err)
		return
	}

	data, err := os.ReadFile("comment.txt")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println(string(data))
}
