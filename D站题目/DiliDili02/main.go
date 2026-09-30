package main

import (
	"fmt"
)

func main() {
	playCounts := []int{120, 350, 400, 200, 150, 500, 280}
	total := 0
	for _, playCount := range playCounts {
		total += playCount
	}

	fmt.Printf("总播放量为%d\n", total)
}
