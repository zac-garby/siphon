package server

import (
	"fmt"
	"net/http"

	"github.com/Zac-Garby/db/db"
	"github.com/gorilla/mux"
)

// A Server listens for query requests over HTTP and manages a database instance.
type Server struct {
	Addr     string
	Database *db.DB
}

// NewServer makes a new server, initialising a database from the schema string.
func NewServer(addr, schema string) (*Server, error) {
	sch := &db.Schema{}
	if err := db.SchemaParser.ParseString(schema, sch); err != nil {
		return nil, err
	}

	d, err := db.MakeDB(sch)
	if err != nil {
		return nil, err
	}

	return &Server{
		Addr:     addr,
		Database: d,
	}, nil
}

// Listen starts listening on the given address.
func (s *Server) Listen() error {
	r := mux.NewRouter()
	r.HandleFunc("/json", s.handleJSONSelector)

	return http.ListenAndServe(s.Addr, r)
}

func (s *Server) handleJSONSelector(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(r.Form["selector"]) != 1 {
		http.Error(w, "only one form value expected for the selector", http.StatusInternalServerError)
		return
	}

	res, status := s.Database.QueryString(r.Form["selector"][0])
	if status != db.StatusOK {
		errorMessage(w, status)
		return
	}

	w.Header().Set("Content-Type", "text/json")
	fmt.Fprint(w, res.JSON())
}

func errorMessage(w http.ResponseWriter, status string) {
	w.Header().Set("Content-Type", "text/json")
	var msg string
	switch status {
	case db.StatusError:
		msg = "unknown error"
	case db.StatusIndex:
		msg = "invalid index or key"
	case db.StatusNOOP:
		msg = "invalid operation"
	case db.StatusNoType:
		msg = "unknown type"
	case db.StatusType:
		msg = "invalid type"
	default:
		msg = "unknown status code: " + status
	}
	http.Error(w, fmt.Sprintf(`{"error": "%s"}`, msg), http.StatusInternalServerError)
}
