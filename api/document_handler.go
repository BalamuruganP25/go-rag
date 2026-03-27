package api

import (
	"encoding/json"
	"net/http"

	"go-rag/rag"
	"go-rag/storage"
)

type DocumentRequest struct {
	Text string `json:"text"`
}

type DocumentResponse struct {
	Message string `json:"message"`
}

func DocumentHandler(w http.ResponseWriter, r *http.Request) {

	var req DocumentRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	chunks := rag.ChunkDocument(req.Text, 50)

	embedder := rag.NewEmbedder()
	db := storage.NewVectorDB()

	for _, chunk := range chunks {

		vector, err := embedder.Embed(chunk.Text)
		if err != nil {
			http.Error(w, "embedding failed", http.StatusInternalServerError)
			return
		}

		err = db.Insert(int64(chunk.ID), vector, chunk.Text)
		if err != nil {
			http.Error(w, "vector insert failed", http.StatusInternalServerError)
			return
		}
	}

	res := DocumentResponse{
		Message: "document ingested successfully",
	}

	json.NewEncoder(w).Encode(res)
}
