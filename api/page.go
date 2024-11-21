package api

import (
	"github.com/flosch/pongo2/v6"
	"github.com/saintbyte/far-away/pkg/db"
	"github.com/saintbyte/far-away/pkg/templates"
	"html"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	err := db.ConnectPG()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var tplExample = pongo2.Must(pongo2.FromString(templates.PageTemplate))
	err = tplExample.ExecuteWriter(
		pongo2.Context{
			"query":       r.FormValue("query"),
			"title":       html.EscapeString("Home"),
			"description": "Write your history here",
		},
		w,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
