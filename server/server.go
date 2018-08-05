package server

import (
	"encoding/json"
	"net/http"

	"github.com/Zac-Garby/db/db"
	"github.com/gorilla/mux"
)

// A Server listens for query requests over HTTP and manages a database instance.
type Server struct {
	Addr string
}

// Listen starts listening on the given address.
func (s *Server) Listen() error {
	r := mux.NewRouter()
	r.HandleFunc("/json", handleJSONSelector)

	return http.ListenAndServe(s.Addr, r)
}

func handleJSONSelector(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(r.Form["selector"]) != 1 {
		http.Error(w, "only one form value expected for the selector", http.StatusInternalServerError)
		return
	}

	var (
		selector = r.Form["selector"][0]
		parsed   = &db.Selector{}
		err      = db.SelectorParser.ParseString(selector, parsed)
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/json")
	w.Write(bytes)
}
