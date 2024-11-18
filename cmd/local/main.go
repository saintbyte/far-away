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
	mux.HandleFunc("/api/healthcheck", handler.Healthcheck)

	mux.Handle("/i/", http.StripPrefix("/i/", http.FileServer(http.Dir("./public/i"))))
	mux.Handle("/favicon/", http.StripPrefix("/favicon/", http.FileServer(http.Dir("./public/favicon"))))

	mux.HandleFunc("/{page}", handler.Page)
	mux.HandleFunc("/", handler.Index)

	err = http.ListenAndServe(":"+utils.GetPort(), utils.LogRequest(mux))
	if err != nil {
		slog.Error("ListenAndServe problem:", err)
	}
}
