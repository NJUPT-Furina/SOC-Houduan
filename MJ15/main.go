package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		fmt.Fprintf(w, "视频编号：%s\n", id)
	})

	fmt.Println("HTTP 服务启动：http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
