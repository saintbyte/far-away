package main

import (
	"github.com/joho/godotenv"
	handler "github.com/saintbyte/far-away/api"
	"github.com/saintbyte/far-away/pkg/utils"
	"log/slog"
	"net/http"
)

func main() {
	// ENV
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/healthcheck", handler.Healthcheck)
	mux.HandleFunc("/api/setup", handler.Setup)
	mux.HandleFunc("/api/page", handler.Page)
	mux.HandleFunc("/api/abuse", handler.Abuse)
	mux.HandleFunc("/api/save", handler.Save)

	mux.Handle("/i/", http.StripPrefix("/i/", http.FileServer(http.Dir("./public/i"))))
	mux.Handle("/favicon/", http.StripPrefix("/favicon/", http.FileServer(http.Dir("./public/favicon"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets"))))

	mux.HandleFunc("/{page}", handler.Page)
	mux.HandleFunc("/", handler.Index)

	err = http.ListenAndServe(":"+utils.GetPort(), utils.LogRequest(mux))
	if err != nil {
		slog.Error("ListenAndServe problem:", err)
	}
}
