package utils

import (
	"log"
	"net/http"
	"os"
)

func GetPort() string {
	// Получаем порт на котором надо запуститься
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func LogRequest(handler http.Handler) http.Handler {
	// Милдварь для логов
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
