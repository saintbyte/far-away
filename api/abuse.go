package api

import (
	"fmt"
	"net/http"
)

func Abuse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Index")
	fmt.Fprintf(w, "<br />")
}
