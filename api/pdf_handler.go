package api

import (
	"go-rag/rag"
	"go-rag/storage"
	"io"
	"net/http"
	"os"

	"github.com/ledongthuc/pdf"
)

func UploadPDFHandler(w http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file upload failed", http.StatusBadRequest)
		return
	}

	defer file.Close()

	tmpFile, err := os.CreateTemp("", "upload-*.pdf")
	if err != nil {
		http.Error(w, "temp file error", http.StatusInternalServerError)
		return
	}

	defer os.Remove(tmpFile.Name())

	io.Copy(tmpFile, file)

	text, err := extractPDFText(tmpFile.Name())
	if err != nil {
		http.Error(w, "pdf parsing failed", http.StatusInternalServerError)
		return
	}

	chunks := rag.ChunkDocument(text, 100)

	embedder := rag.NewEmbedder()
	db := storage.NewVectorDB()

	for _, chunk := range chunks {

		vector, err := embedder.Embed(chunk.Text)
		if err != nil {
			http.Error(w, "embedding failed", http.StatusInternalServerError)
			return
		}

		db.Insert(int64(chunk.ID), vector, chunk.Text)
	}

	w.Write([]byte("PDF ingested successfully"))
}

func extractPDFText(path string) (string, error) {

	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	var text string

	totalPage := r.NumPage()

	for i := 1; i <= totalPage; i++ {

		page := r.Page(i)

		if page.V.IsNull() {
			continue
		}

		content, _ := page.GetPlainText(nil)

		text += content
	}

	return text, nil
}
