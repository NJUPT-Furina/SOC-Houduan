package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Account struct {
	AccountName string `json:"accountname"`
	Name        string `json:"name"`
}

func main() {
	http.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "只支持 POST", http.StatusMethodNotAllowed)
			return
		}

		var a Account

		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, "解析失败", http.StatusBadRequest)
			return
		}

		fmt.Printf("Hello World!   %s\n", a.AccountName)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"Status": "成功"})
	})

	fmt.Println("HTTP 服务启动：http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
