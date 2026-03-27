package storage

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type VectorDB struct {
	BaseURL    string
	Collection string
}

func NewVectorDB() *VectorDB {

	return &VectorDB{
		BaseURL:    "http://localhost:6333",
		Collection: "documents",
	}
}

type Document struct {
	Text string
}

type searchRequest struct {
	Vector []float64 `json:"vector"`
	Limit  int       `json:"limit"`
}

type searchResponse struct {
	Result []struct {
		Payload struct {
			Text string `json:"text"`
		} `json:"payload"`
	} `json:"result"`
}

type point struct {
	ID      int64     `json:"id"`
	Vector  []float64 `json:"vector"`
	Payload payload   `json:"payload"`
}

type payload struct {
	Text string `json:"text"`
}

type insertRequest struct {
	Points []point `json:"points"`
}

func (v *VectorDB) Search(vector []float64) ([]Document, error) {

	reqBody := searchRequest{
		Vector: vector,
		Limit:  3,
	}

	jsonBody, _ := json.Marshal(reqBody)

	url := v.BaseURL + "/collections/" + v.Collection + "/points/search"

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result searchResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var docs []Document

	for _, r := range result.Result {
		docs = append(docs, Document{
			Text: r.Payload.Text,
		})
	}

	return docs, nil
}

func (v *VectorDB) Insert(id int64, vector []float64, text string) error {

	reqBody := insertRequest{
		Points: []point{
			{
				ID:     id,
				Vector: vector,
				Payload: payload{
					Text: text,
				},
			},
		},
	}

	jsonBody, _ := json.Marshal(reqBody)

	url := v.BaseURL + "/collections/" + v.Collection + "/points"

	req, err := http.NewRequest(
		http.MethodPut,
		url,
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}