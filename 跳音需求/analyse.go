package main

import "fmt"

type AnalysisResult struct {
	TotalContents int
	TopContent    Content
	TopScore      float64
	AvgScores     map[string]float64
}

type TrafficAnalyzer struct {
	contents []Content
}

func NewTrafficAnalyzer(contents []Content) *TrafficAnalyzer {
	return &TrafficAnalyzer{
		contents: contents,
	}
}

func (ta *TrafficAnalyzer) Analyze() AnalysisResult {
	result := AnalysisResult{
		TotalContents: len(ta.contents),
		AvgScores:     make(map[string]float64),
	}

	platformTotal := make(map[string]float64)
	platformCount := make(map[string]int)

	for _, c := range ta.contents {
		platform := GetPlatform(c.Platform)
		if platform == nil {
			continue
		}

		score := platform.CalculateTrafficScore(c)

		if score > result.TopScore {
			result.TopScore = score
			result.TopContent = c
		}

		platformTotal[c.Platform] += score
		platformCount[c.Platform]++
	}

	for name, total := range platformTotal {
		count := platformCount[name]
		if count > 0 {
			result.AvgScores[name] = total / float64(count)
		}
	}

	return result
}

func PrintAnalysisReport(result AnalysisResult) {
	fmt.Println("\n========== 流量分析报告 ==========")
	fmt.Printf("总内容数：%d\n", result.TotalContents)

	if result.TotalContents == 0 {
		fmt.Println("暂无内容可供分析")
		fmt.Println("=====================================")
		return
	}

	fmt.Printf("最高分内容：\n")
	fmt.Printf("编号：%d\n", result.TopContent.ID)
	fmt.Printf("标题：%s\n", result.TopContent.Title)
	fmt.Printf("作者：%s\n", result.TopContent.Author)
	fmt.Printf("平台：%s\n", result.TopContent.Platform)
	fmt.Printf("得分：%.2f\n", result.TopScore)

	fmt.Printf("各平台平均分：\n")
	if len(result.AvgScores) == 0 {
		fmt.Println("暂无数据")
	} else {
		for platform, avg := range result.AvgScores {
			fmt.Printf("   %s：%.2f\n", platform, avg)
		}
	}
	fmt.Println("=====================================")
}
