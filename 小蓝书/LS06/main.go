package main

import (
	"fmt"
)

type Note struct {
	title string
	likes int
}

func main() {
	n := Note{title: "JavaStart", likes: 7}
	for i := 0; i < 3; i++ {
		n.likes = addLike(n.likes)
	}
	getInfo(n.title, n.likes)
}

func addLike(like int) int {
	like++
	return like
}

func getInfo(title string, like int) {
	fmt.Printf("%s：%d", title, like)
}
