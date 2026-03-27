package rag

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Generator struct {
	BaseURL string
	Model   string
}

func NewGenerator() *Generator {

	return &Generator{
		BaseURL: "http://localhost:11434",
		Model:   "phi3",
	}
}

func (g *Generator) Generate(context string, question string) (string, error) {

	prompt := "Context:\n" + context + "\n\nQuestion:\n" + question + "\nAnswer:"

	reqBody := map[string]interface{}{
		"model":  g.Model,
		"prompt": prompt,
		"stream": false,
	}

	jsonBody, _ := json.Marshal(reqBody)

	resp, err := http.Post(
		g.BaseURL+"/api/generate",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result struct {
		Response string `json:"response"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Response, nil
}