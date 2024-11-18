package main

import (
	"github.com/joho/godotenv"
	handler "github.com/saintbyte/far-away/api"
	"github.com/saintbyte/far-away/pkg/utils"
	"log/slog"
	"net/http"
	"runtime"
)

func main() {
	// ENV
	err := godotenv.Load()
	runtime.GOMAXPROCS(100)
	if err != nil {
		slog.Error("Error loading .env file")
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/api/healthcheck", handler.Healthcheck)

	mux.HandleFunc("/", handler.Index)

	mux.HandleFunc("/{page:[a-z]+}", handler.Page)

	err = http.ListenAndServe(":"+utils.GetPort(), utils.LogRequest(mux))
	if err != nil {
		slog.Error("ListenAndServe problem:", err)
	}
}
