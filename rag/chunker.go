package rag

import "strings"

type Chunk struct {
	ID   int
	Text string
}

func ChunkDocument(text string, chunkSize int) []Chunk {

	var chunks []Chunk

	words := strings.Fields(text)

	if len(words) == 0 {
		return chunks
	}

	start := 0
	id := 1

	for start < len(words) {

		end := start + chunkSize

		if end > len(words) {
			end = len(words)
		}

		chunkWords := words[start:end]

		chunkText := strings.Join(chunkWords, " ")

		chunks = append(chunks, Chunk{
			ID:   id,
			Text: chunkText,
		})

		id++
		start = end
	}

	return chunks
}