package main

import (
	"fmt"
)

type Note struct {
	likes    int
	saves    int
	comments int
}

func main() {
	notes := [3]Note{
		{likes: 3, saves: 4, comments: 1},
		{likes: 10, saves: 5, comments: 8},
		{likes: 0, saves: 2, comments: 6},
	}

	scores := make([]int, 0, len(notes))
	for _, note := range notes {
		score := calculateScore(note.likes, note.saves, note.comments)
		scores = append(scores, score)
	}
	newScore := rankScore(scores)
	fmt.Println(newScore)
}

func calculateScore(likes int, saves int, comment int) int {
	score := likes*2 + saves*3 + comment
	return score
}

func rankScore(scores []int) []int {
	n := len(scores)
	for i := 0; i < n-1; i++ {
		maxScore := i
		for j := i + 1; j < n; j++ {
			if scores[j] > scores[maxScore] {
				maxScore = j
			}
		}
		scores[i], scores[maxScore] = scores[maxScore], scores[i]
	}
	return scores
}
