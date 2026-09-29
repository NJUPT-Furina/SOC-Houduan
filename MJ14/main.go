package main

import (
	"fmt"
	"sync"
)

type Account struct {
	AccountName string
	Name        string
	password    string
	result      bool
}

func main() {
	accounts := []Account{
		{AccountName: "cqx001", Name: "cqx", password: "26040721cqx", result: true},
		{AccountName: "cqx002", Name: "chenqixxing", password: "26040722cqx", result: true},
		{AccountName: "cqx003", Name: "陈启昕", password: "26040723cqx", result: true},
		{AccountName: "cqx004", Name: "ttdr", password: "26040724cqx", result: true},
		{AccountName: "cqx005", Name: "cqx14138", password: "26040725cqx", result: true},
	}
	count := 0
	var wg sync.WaitGroup
	for _, account := range accounts {
		account := account
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println("正在创建账号：", account.AccountName)
		}()

		count = count + 1
	}
	wg.Wait()
	fmt.Printf("共创建账号：%d个", count)
}
