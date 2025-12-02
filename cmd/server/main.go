package main

import (
	"fmt"
	"net/http"
)

func main() {
	//注册 /ping 路由
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// 简单返回文本 "pong"
		_, _ = w.Write([]byte("pong"))
	})

	fmt.Println("String server on :8080")

	// 启动 HTTP 服务器，监听 8080 端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}
}
