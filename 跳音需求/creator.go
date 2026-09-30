package main

import (
	"fmt"
)

func NewCreator(name string) *Creator {
	return &Creator{
		Name:              name,
		contentCount:      0,
		totalTrafficScore: 0,
		contents:          make([]Content, 0),
	}
}

func (cr *Creator) AddContent(c Content) {
	cr.contents = append(cr.contents, c)
	cr.contentCount++
	platform := GetPlatform(c.Platform)
	if platform != nil {
		cr.totalTrafficScore += platform.CalculateTrafficScore(c)
	}
}

func (cr *Creator) GetCreatorLevel() string {
	if cr.totalTrafficScore >= 500000 {
		return "平台顶流"
	} else if cr.totalTrafficScore >= 100000 {
		return "热门创作者"
	} else if cr.totalTrafficScore >= 10000 {
		return "潜力创作者"
	} else {
		return "新人创作者"
	}
}

func (cr *Creator) GetContentCount() int {
	return cr.contentCount
}

func (cr *Creator) GetTotalTrafficScore() float64 {
	return cr.totalTrafficScore
}

func showCreators() {
	fmt.Println("\n=== 创作者统计 ===")

	if len(creators) == 0 {
		fmt.Println("暂无创作者数据")
		return
	}

	for name, cr := range creators {
		fmt.Printf("创作者：%s\n", name)
		fmt.Printf("发布内容数：%d\n", cr.GetContentCount())
		fmt.Printf("总流量分数：%.2f\n", cr.GetTotalTrafficScore())
		fmt.Printf("创作者等级：%s\n", cr.GetCreatorLevel())
	}
}
