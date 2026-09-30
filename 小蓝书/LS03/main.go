package main

import (
	"fmt"
)

func main() {
	comments := [5]int{0, 18, 21, 30, 20}
	count := 0
	for _, comment := range comments {
		if comment >= 20 {
			count++
		}
	}
	fmt.Println(count)
}
