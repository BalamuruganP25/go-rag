package main

import (
	"log"
	"net/http"

	"go-rag/api"
	"go-rag/rag"
)

func main() {

	retriever := rag.NewRetriever()
	generator := rag.NewGenerator()

	ragService := rag.NewService(retriever, generator)

	handler := &api.Handler{
		RagService: ragService,
	}

	router := api.NewRouter(handler)

	log.Println("Server running on port 8080")

	http.ListenAndServe(":8080", router)
}
