package main

import (
	"fmt"
)

type Information struct {
	Name string
	Note int
	Tag  string
}

func main() {
	i := Information{Name: "ttdr", Note: 3, Tag: "coding"}
	fmt.Printf("%s has %d notes and follows %s.", i.Name, i.Note, i.Tag)
}
