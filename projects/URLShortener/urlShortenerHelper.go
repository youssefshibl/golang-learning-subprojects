package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/save", saveUrl)
	mux.HandleFunc("/", redirect)
}

var urlMap = make(map[string]string)
var hashMap = make(map[string]string)

func saveUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		sendJSON(w, map[string]string{
			"error": "Invalid form data",
		}, http.StatusBadRequest)
		return
	}
	url := r.FormValue("url")
	if url == "" {
		sendJSON(w, map[string]string{
			"error": "Missing 'url' field",
		}, http.StatusBadRequest)
		return
	}
	urlHash := hashUrl(url)
	var code string

	if _, exists := hashMap[urlHash]; !exists {
		code = RandStringRunes(16)
		urlMap[code] = url
		hashMap[urlHash] = code
	} else {
		code = hashMap[urlHash]
	}

	fmt.Println(urlMap)
	sendJSON(w, map[string]string{
		"code": code,
	}, 200)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path // e.g., "/1245351"
	code := strings.TrimPrefix(path, "/")
	if _, exists := urlMap[code]; !exists {
		http.NotFound(w, r)
	}
	http.Redirect(w, r, urlMap[code], http.StatusFound)
}

func sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
