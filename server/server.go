package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Zac-Garby/siphon/db"
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
	r.HandleFunc("/json", s.handleJSON)
	r.HandleFunc("/set", s.handleSet)
	r.HandleFunc("/unset", s.handleUnset)
	r.HandleFunc("/append", s.handleAppend)
	r.HandleFunc("/prepend", s.handlePrepend)
	r.HandleFunc("/key", s.handleKey)
	r.HandleFunc("/empty", s.handleEmpty)

	return http.ListenAndServe(s.Addr, r)
}

func (s *Server) handleJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")

	if r.Method != "GET" {
		errorMessage(w, "only GET is supported for /json")
		return
	}

	if err := r.ParseForm(); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if len(r.Form["selector"]) != 1 {
		errorMessage(w, "only one form value expected for the selector")
		return
	}

	selector, err := url.QueryUnescape(r.Form["selector"][0])
	if err != nil {
		errorMessage(w, "could not unescape selector: "+r.Form["selector"][0])
		return
	}

	res, err := s.Database.QueryString(selector)
	if err != nil {
		errorMessage(w, err.Error())
		return
	}

	fmt.Fprint(w, res.JSON())
}

func (s *Server) handleSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")

	if r.Method != "POST" {
		errorMessage(w, "only POST is supported for /set")
		return
	}

	if err := r.ParseForm(); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if len(r.Form["selector"]) != 1 {
		errorMessage(w, "only one form value expected for the selector")
		return
	}

	selector, err := url.QueryUnescape(r.Form["selector"][0])
	if err != nil {
		errorMessage(w, "could not unescape selector: "+r.Form["selector"][0])
		return
	}

	if r.Body == nil {
		errorMessage(w, "expected a request body")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMessage(w, "could not read request body")
		return
	}

	item, err := s.Database.QueryString(selector)
	if err != nil {
		errorMessage(w, err.Error())
		return
	}

	var val interface{}
	if err := json.Unmarshal(body, &val); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if err = item.Set(val); err != nil {
		errorMessage(w, err.Error())
		return
	}
}

func (s *Server) handleAppend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")

	if r.Method != "POST" {
		errorMessage(w, "only POST is supported for /append")
		return
	}

	if err := r.ParseForm(); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if len(r.Form["selector"]) != 1 {
		errorMessage(w, "only one form value expected for the selector")
		return
	}

	selector, err := url.QueryUnescape(r.Form["selector"][0])
	if err != nil {
		errorMessage(w, "could not unescape selector: "+r.Form["selector"][0])
		return
	}

	if r.Body == nil {
		errorMessage(w, "expected a request body")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMessage(w, "could not read request body")
		return
	}

	item, err := s.Database.QueryString(selector)
	if err != nil {
		errorMessage(w, err.Error())
		return
	}

	var val interface{}
	if err := json.Unmarshal(body, &val); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if err = item.AppendJSON(val); err != nil {
		errorMessage(w, err.Error())
		return
	}
}

func (s *Server) handlePrepend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")

	if r.Method != "POST" {
		errorMessage(w, "only POST is supported for /prepend")
		return
	}

	if err := r.ParseForm(); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if len(r.Form["selector"]) != 1 {
		errorMessage(w, "only one form value expected for the selector")
		return
	}

	selector, err := url.QueryUnescape(r.Form["selector"][0])
	if err != nil {
		errorMessage(w, "could not unescape selector: "+r.Form["selector"][0])
		return
	}

	if r.Body == nil {
		errorMessage(w, "expected a request body")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMessage(w, "could not read request body")
		return
	}

	item, err := s.Database.QueryString(selector)
	if err != nil {
		errorMessage(w, err.Error())
		return
	}

	var val interface{}
	if err := json.Unmarshal(body, &val); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if err = item.PrependJSON(val); err != nil {
		errorMessage(w, err.Error())
		return
	}
}

func (s *Server) handleKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")

	if r.Method != "POST" {
		errorMessage(w, "only POST is supported for /key")
		return
	}

	if err := r.ParseForm(); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if len(r.Form["selector"]) != 1 {
		errorMessage(w, "only one form value expected for the selector")
		return
	}

	selector, err := url.QueryUnescape(r.Form["selector"][0])
	if err != nil {
		errorMessage(w, "could not unescape selector: "+r.Form["selector"][0])
		return
	}

	if r.Body == nil {
		errorMessage(w, "expected a request body")
		return
	}

	item, err := s.Database.QueryString(selector)
	if err != nil {
		errorMessage(w, err.Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMessage(w, "could not read request body")
		return
	}

	data := struct {
		Key   interface{} `json:"key"`
		Value interface{} `json:"value"`
	}{}

	if err := json.Unmarshal(body, &data); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if err := item.SetKeyJSON(data.Key, data.Value); err != nil {
		errorMessage(w, err.Error())
		return
	}
}

func (s *Server) handleUnset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")

	if r.Method != "POST" {
		errorMessage(w, "only POST is supported for /unset")
		return
	}

	if err := r.ParseForm(); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if len(r.Form["selector"]) != 1 {
		errorMessage(w, "only one form value expected for the selector")
		return
	}

	selector, err := url.QueryUnescape(r.Form["selector"][0])
	if err != nil {
		errorMessage(w, "could not unescape selector: "+r.Form["selector"][0])
		return
	}

	if r.Body == nil {
		errorMessage(w, "expected a request body")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMessage(w, "could not read request body")
		return
	}

	item, err := s.Database.QueryString(selector)
	if err != nil {
		errorMessage(w, err.Error())
		return
	}

	var val interface{}
	if err := json.Unmarshal(body, &val); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if err = item.UnsetKeyJSON(val); err != nil {
		errorMessage(w, err.Error())
		return
	}
}

func (s *Server) handleEmpty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")

	if r.Method != "POST" {
		errorMessage(w, "only POST is supported for /empty")
		return
	}

	if err := r.ParseForm(); err != nil {
		errorMessage(w, err.Error())
		return
	}

	if len(r.Form["selector"]) != 1 {
		errorMessage(w, "only one form value expected for the selector")
		return
	}

	selector, err := url.QueryUnescape(r.Form["selector"][0])
	if err != nil {
		errorMessage(w, "could not unescape selector: "+r.Form["selector"][0])
		return
	}

	res, err := s.Database.QueryString(selector)
	if err != nil {
		errorMessage(w, err.Error())
		return
	}

	if err = res.Empty(); err != nil {
		errorMessage(w, err.Error())
		return
	}
}

func errorMessage(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)

	bytes, err := json.Marshal(map[string]string{
		"err": msg,
	})

	if err != nil {
		fmt.Fprint(w, `{"err": "couldn't convert error message to JSON"}`)
		return
	}

	w.Write(bytes)
}
