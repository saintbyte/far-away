package api

import (
	"fmt"
	"github.com/saintbyte/far-away/pkg/db"
	"github.com/saintbyte/far-away/pkg/models"
	"github.com/saintbyte/far-away/pkg/utils"
	"log/slog"
	"net/http"
)

func Setup(w http.ResponseWriter, r *http.Request) {
	err := utils.CheckAccessKey(w, r)
	if err != nil {
		return
	}
	err = db.ConnectPG()
	if err != nil {
		slog.Error("Pg connect error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Database.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	err = db.Database.AutoMigrate(&models.PageDBModel{})
	if err != nil {
		slog.Error("AutoMigrate error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Database.Raw("CREATE UNIQUE INDEX IF NOT EXISTS slug_index_uniq ON page_db_models(slug);")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "If you seem if setup is OK")
}
