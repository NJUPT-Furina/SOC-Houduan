package main

import (
	"fmt"
)

func main() {
	hots := []int{2, 7, 11, 15}
	target := 9
	result := catch(hots, target)
	if result == nil {
		fmt.Println("未找到匹配项！")
	}
	fmt.Println(result)
}

func catch(hots []int, target int) []int {
	for i := 0; i < len(hots); i++ {
		for j := 0; j < len(hots); j++ {
			if hots[i]+hots[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
