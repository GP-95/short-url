package url

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
)

type urlBody struct {
	Url string `json:"url"`
}

type codeBody struct {
	Code string `json:"code"`
}

func RegisterHandlers() {
	http.HandleFunc("POST /shorten", addUrlRoute)
	http.HandleFunc("GET /{hash}", getUrlRoute)
}

func addUrlRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	var body urlBody
	err := decoder.Decode(&body)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	defer r.Body.Close()

	// Ideally we'd use an external validation library or a regex
	_, err = url.ParseRequestURI(body.Url)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("We only accept valid URLs"))
		return
	}

	hash, err := addNewUrl(r.Context(), body.Url)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	res := codeBody{Code: hash}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func getUrlRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hash := r.PathValue("hash")
	if len(hash) < 11 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not find requested URL."))
		return
	}

	url, err := getCodeUrl(r.Context(), hash)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	if len(url) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Url not found."))
		return
	}

	res := urlBody{Url: url}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}
