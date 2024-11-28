package api

import (
	"errors"
	"github.com/flosch/pongo2/v6"
	"github.com/microcosm-cc/bluemonday"
	"github.com/saintbyte/far-away/pkg/db"
	"github.com/saintbyte/far-away/pkg/models"
	"github.com/saintbyte/far-away/pkg/templates"
	"net/http"
)

func getDescription(text string) string {
	// Получить кусок с описание для страницы
	if text == "" {
		return ""
	}
	p := bluemonday.StripTagsPolicy()
	fullText := p.Sanitize(
		text,
	)
	if fullText == "" {
		return ""
	}
	var description string
	if len(fullText) <= templates.MaxDescriptionLength {
		description = fullText
	} else {
		description = fullText[:templates.MaxDescriptionLength-1]
	}
	return description
}

func Page(w http.ResponseWriter, r *http.Request) {
	err := db.ConnectPG()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	pageSlug := r.URL.Query().Get("page")
	if pageSlug == "" {
		http.Error(w, errors.New("empty request").Error(), http.StatusBadRequest)
		return
	}
	dbRecord := models.PageDBModel{}
	result := db.Database.Model(&models.PageDBModel{}).Where("slug = ?", pageSlug).Take(&dbRecord)
	if result.RowsAffected == 0 {
		http.Error(w, errors.New("Not found: "+pageSlug).Error(), http.StatusNotFound)
		return
	}
	var tplExample = pongo2.Must(pongo2.FromString(templates.PageTemplate))
	// --------------------------------------------------

	err = tplExample.ExecuteWriter(
		pongo2.Context{
			"title":       dbRecord.Title,
			"author":      dbRecord.Author,
			"text":        dbRecord.Text,
			"description": getDescription(dbRecord.Text),
		},
		w,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
