package api

import (
	"errors"
	"github.com/flosch/pongo2/v6"
	"github.com/microcosm-cc/bluemonday"
	"github.com/saintbyte/far-away/pkg/db"
	"github.com/saintbyte/far-away/pkg/models"
	"github.com/saintbyte/far-away/pkg/templates"
	"html"
	"net/http"
)

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
	p := bluemonday.StripTagsPolicy()
	fullText := p.Sanitize(
		dbRecord.Text,
	)
	err = tplExample.ExecuteWriter(
		pongo2.Context{
			"title":       html.EscapeString(dbRecord.Title),
			"author":      html.EscapeString(dbRecord.Author),
			"text":        dbRecord.Text,
			"description": fullText[:512],
		},
		w,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
