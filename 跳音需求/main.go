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
	playCount    int
	likeCount    int
	commentCount int
	shareCount   int
	saveCount    int
	tags         string
}

type Creator struct {
	Name              string
	contentCount      int
	totalTrafficScore float64
	contents          []Content
}

var (
	contents = make([]Content, 0)
	creators = make(map[string]*Creator)
	nextID   = 1
)

func main() {
	fmt.Printf("欢迎来到跳音流量观察台\n当前身份：后端组 SOC 实习生\n负责人：yantz")
	for {
		fmt.Println("请选择你要进行的操作：")
		fmt.Println("1. 添加一条内容数据")
		fmt.Println("2. 查看所有内容数据")
		fmt.Println("3. 计算并展示流量等级")
		fmt.Println("4. 搜索指定标题或作者")
		fmt.Println("5.查看创作者统计")
		fmt.Println("6. 查看流量综合分析报告")
		fmt.Println("0. 退出系统")
		reader := bufio.NewReader(os.Stdin)
		input := readString(reader, "请键入操作前的编号进入操作：")
		switch input {
		case "1":
			addContent(reader)
		case "2":
			listContents()
		case "3":
			showTrafficScores()
		case "4":
			searchContents(reader)
		case "5":
			showCreators()
		case "6":
			if len(contents) == 0 {
				fmt.Println("暂无数据，请先添加内容")
			} else {
				analyzer := NewTrafficAnalyzer(contents)
				result := analyzer.Analyze()
				PrintAnalysisReport(result)
			}
		case "0":
			fmt.Println("再见，下班收工")
			return
		default:
			fmt.Println("无效选项，请重新选择")
		}
	}
}
