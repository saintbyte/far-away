package api

import (
	"encoding/json"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/saintbyte/far-away/pkg/db"
	"github.com/saintbyte/far-away/pkg/models"
	"github.com/sym01/htmlsanitizer"
	"log/slog"
	"net/http"
)

func Save(w http.ResponseWriter, r *http.Request) {
	err := db.ConnectPG()
	if err != nil {
		slog.Error("Pg connect error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var requestData models.PageCreateRequest
	err = decoder.Decode(&requestData)
	if err != nil {
		slog.Error("Body decode error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestData.Title == "" {
		slog.Error("Title empty")
		http.Error(w, "Title empty", http.StatusBadRequest)
		return
	}
	if requestData.HTML == "" {
		slog.Error("Title empty")
		http.Error(w, "Html empty", http.StatusBadRequest)
		return
	}
	slug.Lowercase = true
	pageSlug := slug.Make(requestData.Title)

	s := htmlsanitizer.NewHTMLSanitizer()
	s.GlobalAttr = []string{"class"}

	sanitizedHTML, err := s.SanitizeString(requestData.HTML)

	page := models.PageDBModel{
		Slug:         pageSlug,
		Title:        requestData.Title,
		Author:       requestData.Author,
		Text:         sanitizedHTML,
		AccessSecret: "11111",
	}

	result := db.Database.Create(&page)
	if result.Error != nil {
		slog.Error("Database create: ", result.Error)
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{}")
}
