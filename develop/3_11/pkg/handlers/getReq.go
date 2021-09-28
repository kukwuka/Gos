package handlers

import (
	"github.com/kukwuka/Gos/develop/3_11/pkg"
	"net/http"
	"time"
)

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	eventTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.services.EventsForDay(pkg.MyTime{Time: eventTime})
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	NewResultResponse(w, res)
}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	eventTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.services.EventsForWeek(pkg.MyTime{Time: eventTime})
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	NewResultResponse(w, res)
}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	eventTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.services.EventsForMonth(pkg.MyTime{Time: eventTime})
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	NewResultResponse(w, res)
}
