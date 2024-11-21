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
	"strconv"
	"time"
)

func getSlug(Title string) string {
	// Получить не назанятый адрес
	slug.Lowercase = true
	pageSlug := slug.Make(Title)

	dbRecord := models.PageDBModel{}
	result := db.Database.Model(&models.PageDBModel{}).Where("slug = ?", pageSlug).First(&dbRecord)
	if result.RowsAffected == 0 {
		return pageSlug
	}
	t := time.Now()
	pageSlug = pageSlug + "-" + t.Format("2006-01-02")
	result = db.Database.Model(&models.PageDBModel{}).Where("slug = ?", pageSlug).First(&dbRecord)
	if result.RowsAffected == 0 {
		return pageSlug
	}
	var i = 0
	var newSlug = pageSlug
	for {
		i = i + 1
		newSlug = pageSlug + "-" + strconv.Itoa(i)
		result = db.Database.Model(&models.PageDBModel{}).Where("slug = ?", newSlug).First(&dbRecord)
		if result.RowsAffected == 0 {
			return newSlug
		}
	}
	return pageSlug
}

func Save(w http.ResponseWriter, r *http.Request) {
	// Сохранить текст в базу
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
	pageSlug := getSlug(requestData.Title)
	s := htmlsanitizer.NewHTMLSanitizer()
	s.GlobalAttr = []string{"class"}

	sanitizedHTML, err := s.SanitizeString(requestData.HTML)
	if err != nil {
		slog.Error("SanitizeString err: ", err.Error)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
