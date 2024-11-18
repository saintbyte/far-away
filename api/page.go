package api

import (
	"fmt"
	"github.com/saintbyte/far-away/pkg/db"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func Page(w http.ResponseWriter, r *http.Request) {
	err := db.ConnectPG()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Fprintf(w, pair[0])
		fmt.Fprintf(w, "=")
		fmt.Fprintf(w, pair[1])
		fmt.Fprintf(w, "<br />")
	}
	fmt.Fprintf(w, "GOMAXPROCS")
	fmt.Fprintf(w, "=")
	fmt.Fprintf(w, "%v", runtime.NumCPU())
	fmt.Fprintf(w, "<br />")
}
