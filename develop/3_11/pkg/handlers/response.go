package handlers

import (
	"encoding/json"
	"github.com/kukwuka/Gos/develop/3_11/pkg"
	"net/http"
)

type errorResponse struct {
	ErrMsg string `json:"error"`
}

func NewErrorResponse(w http.ResponseWriter, err string, status int) {
	jsonError, _ := json.Marshal(&errorResponse{err})
	http.Error(w, string(jsonError), status)
}

type ResultMsg struct {
	Result []pkg.Event `json:"result"`
}

func NewResultResponse(w http.ResponseWriter, res []pkg.Event) {
	jsonError, _ := json.Marshal(&ResultMsg{res})
	w.WriteHeader(http.StatusOK)
	w.Write(jsonError)
}
