package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("这是服务端aaaaaa")
	_, err := fmt.Fprintf(w, "这是服务端")
	if err != nil {
		return
	}
}

func healtHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("healt check ")
}

func main() {
	// 启动http服务
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healtHanlder)
	err := http.ListenAndServe(":9021", nil)
	if err != nil {
		return
	}
}
