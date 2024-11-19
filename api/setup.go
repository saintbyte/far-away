package api

import (
	"github.com/saintbyte/far-away/pkg/db"
	"github.com/saintbyte/far-away/pkg/models"
	"log/slog"
	"net/http"
	"os"
)

func Setup(w http.ResponseWriter, r *http.Request) {
	accessKey := r.URL.Query().Get("access_key")
	if accessKey == "" {
		http.Error(w, "Access key required", http.StatusBadRequest)
		return
	}
	secretAccessKey, ok := os.LookupEnv("SECRET_ACCESS_KEY")
	if !ok {
		http.Error(w, "Secret access key required", http.StatusBadRequest)
		return
	}
	if secretAccessKey != accessKey {
		http.Error(w, "Access key does not match", http.StatusBadRequest)
	}
	err := db.ConnectPG()
	if err != nil {
		slog.Error("Error connect pg:", err)
		return
	}
	//-------------------------
	db.Database.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.Database.AutoMigrate(&models.PageDBModel{})

}
