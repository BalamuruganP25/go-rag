package api

import (
	"encoding/json"
	"net/http"

	"go-rag/rag"
)

type Handler struct {
	RagService *rag.Service
}

type QuestionRequest struct {
	Question string `json:"question"`
}

type AnswerResponse struct {
	Answer string `json:"answer"`
}

func (h *Handler) AskHandler(w http.ResponseWriter, r *http.Request) {

	var req QuestionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	answer, err := h.RagService.Ask(req.Question)
	if err != nil {
		http.Error(w, "processing error", http.StatusInternalServerError)
		return
	}

	res := AnswerResponse{
		Answer: answer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}