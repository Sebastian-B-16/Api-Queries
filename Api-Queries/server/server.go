package server

import (
	"Api-Queries/scripts"
	"Api-Queries/validation"
	"encoding/json"
	"net/http"
)

type Handler struct {
	mux *http.ServeMux
}

func New(s *http.ServeMux) *Handler {
	h := Handler{s}
	h.registerRoutes()
	return &h
}

//registerRoutes for all http endpoints
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/how_much", h.getPay)
	h.mux.HandleFunc("/list_how_many", h.getYearlyPay)
}

func (h *Handler) getPay(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		status, day := validation.Query(r)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode("Not a valid Date, please choose a date between 1 and 31")
			break
		} else {
			enc := json.NewEncoder(w)
			enc.SetIndent("", "\t")
			enc.Encode(scripts.NextPay(day))
		}
	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			break
		}
	}
}

func (h *Handler) getYearlyPay(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		status, day := validation.Query(r)
		if status == http.StatusBadRequest {
			json.NewEncoder(w).Encode("Not a valid Date, please choose a date between 1 and 31")
			break
		} else {
			enc := json.NewEncoder(w)
			enc.SetIndent("", "\t")
			enc.Encode(scripts.YearlyPay(day))
		}
	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			break
		}
	}
}
