package rag

type Service struct {
	retriever *Retriever
	generator *Generator
}

func NewService(r *Retriever, g *Generator) *Service {
	return &Service{
		retriever: r,
		generator: g,
	}
}

func (s *Service) Ask(question string) (string, error) {

	context, err := s.retriever.Search(question)
	if err != nil {
		return "", err
	}

	answer, err := s.generator.Generate(context, question)
	if err != nil {
		return "", err
	}

	return answer, nil
}
