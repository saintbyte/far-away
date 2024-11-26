package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/saintbyte/far-away/pkg/db"
	"github.com/saintbyte/far-away/pkg/models"
	"github.com/sym01/htmlsanitizer"
	"log/slog"
	"net/http"
	"time"
)

func Update(w http.ResponseWriter, r *http.Request) {
	// Сохранить текст в базу
	err := db.ConnectPG()
	if err != nil {
		slog.Error("Pg connect error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var requestData models.PageUpdateRequest
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
	if requestData.Secret == "" {
		slog.Error("Secret empty")
		http.Error(w, "Html empty", http.StatusBadRequest)
		return
	}
	s := htmlsanitizer.NewHTMLSanitizer()
	s.GlobalAttr = []string{"class"}
	secret := randToken()
	sanitizedHTML, err := s.SanitizeString(requestData.HTML)
	if err != nil {
		slog.Error("SanitizeString err: ", err.Error)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	page := models.PageDBModel{
		Title:        requestData.Title,
		Author:       requestData.Author,
		Text:         sanitizedHTML,
		AccessSecret: secret,
	}
	var dbRecord = models.PageDBModel{}
	db.Database.Model(
		&models.PageDBModel{},
	).Where(
		"slug = ? AND secret = ?", requestData.Slug, requestData.Secret,
	).Take(&dbRecord)
	result := db.Database.Model(&dbRecord).Updates(
		models.PageDBModel{
			Title:     requestData.Title,
			Author:    requestData.Author,
			Text:      requestData.HTML,
			UpdatedAt: time.Now(),
		},
	)
	if result.Error != nil {
		slog.Error("Update Error: ", result.Error)
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
	if result.RowsAffected == 0 {
		updateErr := errors.New("wrong slug or secret")
		slog.Error("Update Error: ", updateErr)
		http.Error(w, updateErr.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	resultData := &models.PageUpdateResponse{
		Slug: page.Slug,
	}
	b, err := json.Marshal(resultData)
	if err != nil {
		slog.Error("to json error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(b))
}
