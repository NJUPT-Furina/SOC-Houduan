package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Video struct {
	Title     string
	Blogger   string
	PlayCount int
	LikeCount int
}

func main() {
	v := Video{Title: "Python后端开发入门", Blogger: "D站技术UP主", PlayCount: 10000, LikeCount: 500}
	fmt.Printf("视频标题：%s\nUP主：%s\n播放量：%d\n点赞数：%d\n", v.Title, v.Blogger, v.PlayCount, v.LikeCount)

	v.LikeCount = likeInput(v.LikeCount)
	fmt.Printf("操作成功！当前点赞量：%d", v.LikeCount)

}

func likeInput(data int) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("点赞请输入1，取消请输入任意内容")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "1" {
		return data + 1
	}
	return data
}
