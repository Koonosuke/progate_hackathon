package usecase

type EmbeddingService interface {
	GetEmbedding(text string) ([]float32, error)
}
