package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Request struct {
	URL string `json:"url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if !IsValidURL(req.URL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	code := GenerateCode(req.URL)
	SaveURL(code, req.URL)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"code": code})
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	url := GetURL(code)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
