package rag

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Embedder struct {
	BaseURL string
	Model   string
}

func NewEmbedder() *Embedder {
	return &Embedder{
		BaseURL: "http://localhost:11434",
		Model:   "phi3",
	}
}

func (e *Embedder) Embed(text string) ([]float64, error) {

	reqBody := map[string]interface{}{
		"model": e.Model,
		"input": text,
	}

	jsonBody, _ := json.Marshal(reqBody)

	resp, err := http.Post(
		e.BaseURL+"/api/embeddings",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result struct {
		Embedding []float64 `json:"embedding"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Embedding, nil
}