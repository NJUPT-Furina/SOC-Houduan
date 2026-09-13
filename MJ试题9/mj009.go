package main

import "fmt"

type Count struct {
	CountName string
	Name      string
	Password  string
	Status    string
	Number    int
}

func main() {
	counts := []Count{{
		CountName: "B26040721",
		Name:      "cqx001",
		Password:  "26040721cqx",
		Status:    "normal",
		Number:    001,
	},
		{
			CountName: "B26040722",
			Name:      "cqx002",
			Password:  "26040722cqx",
			Status:    "normal",
			Number:    002,
		},
		{
			CountName: "B26040723",
			Name:      "cqx003",
			Password:  "26040723cqx",
			Status:    "normal",
			Number:    003,
		},
	}
	p := &counts[1]
	p.Status = "active"
	for _, count := range counts {
		fmt.Printf("账号名：%s\n帐号状态：%s\n", count.CountName, count.Status)
	}
}
