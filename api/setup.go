package api

import (
	"github.com/saintbyte/far-away/pkg/db"
	"github.com/saintbyte/far-away/pkg/models"
	"log/slog"
	"net/http"
)

func Setup(w http.ResponseWriter, r *http.Request) {
	err := db.ConnectPG()
	if err != nil {
		slog.Error("Error connect pg:", err)
		return
	}
	//-------------------------
	db.Database.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.Database.AutoMigrate(&models.PageDBModel{})

}
