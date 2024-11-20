package utils

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
)

func CheckAccessKey(w http.ResponseWriter, r *http.Request) error {
	accessKey := r.URL.Query().Get("access_key")
	if accessKey == "" {
		http.Error(w, "Access key required", http.StatusBadRequest)
		slog.Error("Access key required")
		return errors.New("access key required")
	}
	secretAccessKey, ok := os.LookupEnv("SECRET_ACCESS_KEY")
	if !ok {
		http.Error(w, "Secret access key required", http.StatusBadRequest)
		slog.Error("Secret access key required")
		return errors.New("secret access key required")
	}
	if secretAccessKey == "" {
		http.Error(w, "Secret access key empty", http.StatusBadRequest)
		slog.Error("Secret access key empty")
		return errors.New("secret access key empty")
	}
	if secretAccessKey != accessKey {
		http.Error(w, "Access key does not match", http.StatusBadRequest)
		slog.Error("Access key does not match")
		return errors.New("access key does not match")
	}
	return nil
}
