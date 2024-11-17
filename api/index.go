package api

import (
	"github.com/flosch/pongo2/v6"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tplExample = pongo2.Must(pongo2.FromFile("templates/index.html"))
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
