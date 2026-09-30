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
	content.ID = nextID
	nextID++
	content.Title = readString(reader, "请输入标题：")
	content.Author = readString(reader, "请输入作者：")
	var platformName string
	for {
		platformName = readString(reader, "请输入平台（d站/跳音/小蓝书/慢脚）：")
		if GetPlatform(platformName) != nil {
			break
		}
		fmt.Println("❌ 平台名无效，请重新输入")
	}
	content.playCount = readInt(reader, "请输入播放量/阅读量：")
	content.likeCount = readInt(reader, "请输入点赞数：")
	content.commentCount = readInt(reader, "请输入评论数：")
	content.shareCount = readInt(reader, "请输入转发数：")
	content.saveCount = readInt(reader, "请输入收藏数：")
	content.tags = readString(reader, "请输入标签（用逗号分隔）：")

	fmt.Println("\n=== 数据添加成功 ===")
	fmt.Printf("编号：%d\n", content.ID)
	fmt.Printf("标题：%s\n", content.Title)
	fmt.Printf("作者：%s\n", content.Author)
	fmt.Printf("平台：%s\n", platformName)
	fmt.Printf("播放量：%d\n", content.playCount)
	fmt.Printf("点赞数：%d\n", content.likeCount)
	fmt.Printf("评论数：%d\n", content.commentCount)
	fmt.Printf("转发数：%d\n", content.shareCount)
	fmt.Printf("收藏数：%d\n", content.saveCount)
	fmt.Printf("标签：%s\n", content.tags)

	contents = append(contents, content)

	if _, exists := creators[content.Author]; !exists {
		creators[content.Author] = NewCreator(content.Author)
	}

	creators[content.Author].AddContent(content)
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
		fmt.Printf("播放量：%d\n", c.playCount)
		fmt.Printf("点赞数：%d\n", c.likeCount)
		fmt.Printf("评论数：%d\n", c.commentCount)
		fmt.Printf("转发数：%d\n", c.shareCount)
		fmt.Printf("收藏数：%d\n", c.saveCount)
		fmt.Printf("标签：%s\n", c.tags)
	}

	fmt.Printf("\n共 %d 条数据\n", len(contents))
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
		platform := GetPlatform(c.Platform)
		if platform == nil {
			fmt.Printf("第 %d 条：未知平台 %s\n", i+1, c.Platform)
			continue
		}

		score := platform.CalculateTrafficScore(c)
		level := getTrafficLevel(score)

		fmt.Printf("\n--- 第 %d 条 ---\n", i+1)
		fmt.Printf("标题：%s\n", c.Title)
		fmt.Printf("作者：%s\n", c.Author)
		fmt.Printf("平台：%s\n", platform.GetPlatformName())
		fmt.Printf("流量分数：%.2f\n", score)
		fmt.Printf("流量等级：%s\n", level)
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
