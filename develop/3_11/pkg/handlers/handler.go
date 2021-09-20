package handlers

import (
	"11/pkg/services"
	"fmt"
	"net/http"
	"time"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := &http.ServeMux{}

	mux.HandleFunc("/create_event", logReq(http.HandlerFunc(h.createEvent)))
	mux.HandleFunc("/update_event", logReq(http.HandlerFunc(h.updateEvent)))
	mux.HandleFunc("/delete_event", logReq(http.HandlerFunc(h.deleteEvent)))
	mux.HandleFunc("/events_for_day", logReq(http.HandlerFunc(h.eventsForDay)))
	mux.HandleFunc("/events_for_week", logReq(http.HandlerFunc(h.eventsForWeek)))
	mux.HandleFunc("/events_for_month", logReq(http.HandlerFunc(h.eventsForMonth)))

	return mux
}

func logReq(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var logStr string
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)
		logStr += "Request: " + r.RequestURI + " Time elapsed: " + fmt.Sprint(duration)
		fmt.Println(logStr)
	}
}
