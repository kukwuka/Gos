package handlers

import (
	"11/pkg"
	"encoding/json"
	"net/http"
)

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	req := h.checkJson(w, r)
	if req == nil {
		return
	}
	err := h.services.CreateEvent(req.EventId, req.UserId, req.Date, req.Title, req.Description)
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	NewResultResponse(w, pkg.Events)
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	req := h.checkJson(w, r)
	if req == nil {
		return
	}
	err := h.services.UpdateEvent(req.EventId, req.UserId, req.Date, req.Title, req.Description)
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	NewResultResponse(w, pkg.Events)
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	req := h.checkJson(w, r)
	if req == nil {
		return
	}
	err := h.services.DeleteEvent(req.EventId)
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	NewResultResponse(w, pkg.Events)
}

func (h *Handler) checkJson(w http.ResponseWriter, r *http.Request) *pkg.Event {
	var req pkg.Event
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		NewErrorResponse(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	return &req
}
