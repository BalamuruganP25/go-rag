package rag

import "go-rag/storage"

type Retriever struct {
	vectorDB *storage.VectorDB
	embedder *Embedder
}

func NewRetriever() *Retriever {

	db := storage.NewVectorDB()
	embedder := NewEmbedder()

	return &Retriever{
		vectorDB: db,
		embedder: embedder,
	}
}

func (r *Retriever) Search(question string) (string, error) {

	vector, err := r.embedder.Embed(question)
	if err != nil {
		return "", err
	}

	results, err := r.vectorDB.Search(vector)
	if err != nil {
		return "", err
	}

	context := ""

	for _, doc := range results {
		context += doc.Text + "\n"
	}

	return context, nil
}
