package main

import (
	"fmt"
)

func main() {
	likes := 10
	saves := 5
	score := calculateScore(likes, saves)
	fmt.Println(score)
}

func calculateScore(likes int, saves int) int {
	score := likes*2 + saves*3
	return score
}
