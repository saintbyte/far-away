package api

import (
	"github.com/flosch/pongo2/v6"
	"github.com/saintbyte/far-away/pkg/templates"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tplExample = pongo2.Must(pongo2.FromString(templates.IndexTemplate))
	err := tplExample.ExecuteWriter(
		pongo2.Context{
			"query":       r.FormValue("query"),
			"title":       "Home",
			"description": "Write your history here",
		},
		w,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
