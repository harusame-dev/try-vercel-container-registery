package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Vercel はコンテナの 80 番ポートへルーティングし、プロジェクトの環境変数 PORT で上書きできる
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("hello"))
	})

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
