package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Account struct {
	AccountName string
	Name        string
}

func main() {
	account := Account{AccountName: "cqx001", Name: "陈启昕"}

	body, _ := json.Marshal(account)
	resp, err := http.Post(
		"http://127.0.0.1:8080/account",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}
	defer resp.Body.Close()

	result, _ := io.ReadAll(resp.Body)
	fmt.Println("服务器响应：", string(result))
}
