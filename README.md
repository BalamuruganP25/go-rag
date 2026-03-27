# 🚀 Go-RAG

Go-RAG is a lightweight **Retrieval Augmented Generation (RAG)** system built using Go.
It allows users to upload documents, store embeddings in a vector database, and ask questions powered by a local LLM.

---

## 🧠 Architecture

```
User
 ↓
Go API (Chi Router)
 ↓
RAG Service
 ├─ Retriever → Qdrant
 └─ Generator → Ollama (Phi-3)
 ↓
Answer
```

---

## ⚙️ Tech Stack

* Go (Backend API)
* Chi Router
* Qdrant (Vector Database)
* Ollama (LLM Runtime)
* Phi-3 (LLM Model)
* Docker & Docker Compose

---

## 📦 Features

* ✅ Document ingestion (text)
* ✅ PDF ingestion
* ✅ Chunking pipeline
* ✅ Embedding generation
* ✅ Vector search (semantic retrieval)
* ✅ RAG-based question answering
* ✅ Fully containerized setup

---

## 📁 Project Structure

```
go-rag/
 ├── main.go
 ├── api/
 │    ├── router.go
 │    ├── handler.go
 │    ├── document_handler.go
 │    └── pdf_handler.go
 ├── rag/
 │    ├── service.go
 │    ├── retriever.go
 │    ├── generator.go
 │    ├── embedder.go
 │    └── chunker.go
 ├── storage/
 │    └── vectordb.go
 ├── Dockerfile
 └── docker-compose.yml
```

---

## 🚀 Getting Started

### 1️⃣ Clone Repository

```
git clone <your-repo-url>
cd go-rag
```

---

### 2️⃣ Run with Docker

```
docker compose up --build
```

Services:

* API → http://localhost:8080
* Qdrant → http://localhost:6333
* Ollama → http://localhost:11434

---

## 🤖 LLM Setup (Ollama + Phi-3)

This project uses Ollama to run the Phi-3 model locally.

### 🔽 First-Time Setup

* The Phi-3 model (~2–4GB) will be downloaded automatically on the first request.
* This may take **2–10 minutes**, depending on your internet speed.

During this time, API responses may be slow — this is expected.

---

### ⚡ Optional: Pre-download Model

```
docker exec -it ollama ollama pull phi3
```

---

### 🔍 Verify Model Installation

```
curl http://localhost:11434/api/tags
```

Expected output:

```
{
  "models": [
    {
      "name": "phi3"
    }
  ]
}
```

---

### ⏳ Notes

* Model is stored in Docker volume (`ollama_data`)
* Download happens only once
* Subsequent responses are fast (1–3 seconds)

---

### 🚨 Troubleshooting

```
docker ps
docker logs -f ollama
docker exec -it ollama ollama pull phi3
```

---

## 📄 API Endpoints

### Health Check

```
GET /api/v1/health
```

---

### Upload Document

```
POST /api/v1/documents
```

Example:

```
curl -X POST http://localhost:8080/api/v1/documents \
-H "Content-Type: application/json" \
-d '{
"text":"Go was developed by Robert Griesemer, Rob Pike, and Ken Thompson."
}'
```

---

### Upload PDF

```
POST /api/v1/documents/pdf
```

```
curl -X POST http://localhost:8080/api/v1/documents/pdf \
-F "file=@sample.pdf"
```

---

### Ask Question

```
POST /api/v1/ask
```

Example:

```
curl -X POST http://localhost:8080/api/v1/ask \
-H "Content-Type: application/json" \
-d '{"question":"Who developed Go?"}'
```

---

## 🔄 RAG Pipeline

```
Document
 ↓
Chunking
 ↓
Embedding
 ↓
Store in Qdrant
 ↓
User Question
 ↓
Vector Search
 ↓
Retrieve Context
 ↓
LLM (Phi-3 via Ollama)
 ↓
Answer
```

---

## 🧪 Example

### Input Document

```
Go was developed by Robert Griesemer, Rob Pike, and Ken Thompson.
```

### Question

```
Who developed Go?
```

### Output

```
Robert Griesemer, Rob Pike, and Ken Thompson developed Go.
```

---

## ⚠️ Notes

* First LLM request may take time (model download)
* Ensure Docker is running
* Inside Docker:

  * Ollama → http://ollama:11434
  * Qdrant → http://qdrant:6333

---

## 🚀 Future Improvements

* Metadata support (document_id, source, page)
* Chunk overlap for better retrieval
* Top-K retrieval
* Streaming responses
* Caching (Redis)
* Authentication & rate limiting

---

## 💡 Learning Outcomes

This project demonstrates:

* RAG architecture design
* Vector search systems
* LLM integration
* Go backend development
* Docker-based microservices

---

## 👨‍💻 Author

Bala — Go Developer | AI Backend Enthusiast

---
