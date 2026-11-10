package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type Server struct {
	store   Store
	baseURL string
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func validURL(raw string) bool {
	u, err := url.ParseRequestURI(raw)
	return err == nil && (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
}

func (s *Server) createLink(w http.ResponseWriter, r *http.Request) {
	var in struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if !validURL(in.URL) {
		writeError(w, http.StatusBadRequest, "url must be a valid http or https URL")
		return
	}
	link, err := s.store.Create(in.URL)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create link")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{
		"code":      link.Code,
		"short_url": s.baseURL + "/" + link.Code,
	})
}

func (s *Server) redirect(w http.ResponseWriter, r *http.Request) {
	link, err := s.store.Hit(r.PathValue("code"))
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "link not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "server error")
		return
	}
	http.Redirect(w, r, link.URL, http.StatusFound)
}

func (s *Server) getLink(w http.ResponseWriter, r *http.Request) {
	link, err := s.store.Get(r.PathValue("code"))
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "link not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "server error")
		return
	}
	writeJSON(w, http.StatusOK, link)
}

func (s *Server) deleteLink(w http.ResponseWriter, r *http.Request) {
	err := s.store.Delete(r.PathValue("code"))
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "link not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}