package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func readString(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInt(reader *bufio.Reader, prompt string) int {
	for {
		str := readString(reader, prompt)
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("输入格式错误，请输入数字！")
			continue
		}
		return num
	}
}

func addContent(reader *bufio.Reader) {
	fmt.Println("=== 添加内容数据 ===")

	var content Content
	content.ID = readInt(reader, "请输入内容编号：")
	content.Title = readString(reader, "请输入标题：")
	content.Author = readString(reader, "请输入作者：")
	content.Platform = readString(reader, "请输入平台（d站/跳音/小蓝书/慢脚）：")
	content.PlayCount = readInt(reader, "请输入播放量/阅读量：")
	content.LikeCount = readInt(reader, "请输入点赞数：")
	content.CommentCount = readInt(reader, "请输入评论数：")
	content.ShareCount = readInt(reader, "请输入转发数：")
	content.SaveCount = readInt(reader, "请输入收藏数：")
	content.Tags = readString(reader, "请输入标签（用逗号分隔）：")
	fmt.Println("\n=== 数据添加成功 ===")
	fmt.Printf("编号：%d\n", content.ID)
	fmt.Printf("标题：%s\n", content.Title)
	fmt.Printf("作者：%s\n", content.Author)
	fmt.Printf("平台：%s\n", content.Platform)
	fmt.Printf("播放量：%d\n", content.PlayCount)
	fmt.Printf("点赞数：%d\n", content.LikeCount)
	fmt.Printf("评论数：%d\n", content.CommentCount)
	fmt.Printf("转发数：%d\n", content.ShareCount)
	fmt.Printf("收藏数：%d\n", content.SaveCount)
	fmt.Printf("标签：%s\n", content.Tags)
	contents = append(contents, content)
}

func listContents() {
	fmt.Println("\n=== 已写入的内容数据 ===")

	if len(contents) == 0 {
		fmt.Println("当前还没有内容数据，运营同学还没开始发疯。")
		return
	}

	for i, c := range contents {
		fmt.Printf("\n--- 第 %d 条 ---\n", i+1)
		fmt.Printf("编号：%d\n", c.ID)
		fmt.Printf("标题：%s\n", c.Title)
		fmt.Printf("作者：%s\n", c.Author)
		fmt.Printf("平台：%s\n", c.Platform)
		fmt.Printf("播放量：%d\n", c.PlayCount)
		fmt.Printf("点赞数：%d\n", c.LikeCount)
		fmt.Printf("评论数：%d\n", c.CommentCount)
		fmt.Printf("转发数：%d\n", c.ShareCount)
		fmt.Printf("收藏数：%d\n", c.SaveCount)
		fmt.Printf("标签：%s\n", c.Tags)
	}

	fmt.Printf("\n共 %d 条数据\n", len(contents))
}

func calculateTrafficScore(c Content) float64 {
	return float64(c.PlayCount)*0.4 + float64(c.LikeCount)*2 + float64(c.CommentCount)*3 + float64(c.ShareCount)*4 + float64(c.SaveCount)*5
}

func getTrafficLevel(score float64) string {
	if score >= 200000 {
		return "爆款候选"
	} else if score >= 50000 {
		return "大爆预备"
	} else if score >= 10000 {
		return "小爆一下"
	} else if score >= 1000 {
		return "有点水花"
	} else {
		return "无人问津"
	}
}

func showTrafficScores() {
	fmt.Println("\n=== 所有内容的流量分数与等级 ===")

	if len(contents) == 0 {
		fmt.Println("暂无数据，请先添加内容")
		return
	}

	for i, c := range contents {
		score := calculateTrafficScore(c)
		level := getTrafficLevel(score)

		fmt.Printf("\n--- 第 %d 条 ---\n", i+1)
		fmt.Printf("标题：%s\n", c.Title)
		fmt.Printf("作者：%s\n", c.Author)
		fmt.Printf("平台：%s\n", c.Platform)
		fmt.Printf("🔥 流量分数：%.2f\n", score)
		fmt.Printf("🏆 流量等级：%s\n", level)
	}

	fmt.Printf("\n共 %d 条内容\n", len(contents))
}

func searchContents(reader *bufio.Reader) {
	fmt.Println("\n=== 搜索标题或作者 ===")

	if len(contents) == 0 {
		fmt.Println("暂无数据，请先添加内容")
		return
	}

	keyword := readString(reader, "请输入搜索关键词：")
	if keyword == "" {
		fmt.Println("关键词不能为空")
		return
	}

	found := false
	fmt.Println("\n=== 搜索结果 ===")

	for _, c := range contents {
		if strings.Contains(c.Title, keyword) || strings.Contains(c.Author, keyword) {
			fmt.Printf("标题：%s | 作者：%s\n", c.Title, c.Author)
			found = true
		}
	}

	if !found {
		fmt.Println("没有找到匹配的内容")
	}
}
