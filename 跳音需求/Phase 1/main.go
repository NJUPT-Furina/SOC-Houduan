package main

import (
	"bufio"
	"fmt"
	"os"
)

type Content struct {
	ID           int
	Title        string
	Author       string
	Platform     string
	PlayCount    int
	LikeCount    int
	CommentCount int
	ShareCount   int
	SaveCount    int
	Tags         string
}

var contents []Content

func main() {
	fmt.Printf("欢迎来到跳音流量观察台\n当前身份：后端组 SOC 实习生\n负责人：yantz")
	for {
		fmt.Println("请选择你要进行的操作：")
		fmt.Println("1. 添加一条内容数据")
		fmt.Println("2. 查看所有内容数据")
		fmt.Println("3. 计算并展示流量等级")
		fmt.Println("4. 搜索指定标题或作者")
		fmt.Println("0. 退出系统")
		reader := bufio.NewReader(os.Stdin)
		input := readString(reader, "请键入操作前的编号进入操作：")
		if input == "1" {
			addContent(reader)
		} else if input == "2" {
			listContents()
		} else if input == "3" {
			showTrafficScores()
		} else if input == "4" {
			searchContents(reader)
		} else if input == "0" {
			fmt.Println("再见，下班收工")
			return
		} else {
			fmt.Println("无效选项，请重新选择")
		}
	}
}
