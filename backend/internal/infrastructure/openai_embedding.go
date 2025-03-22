package infrastructure

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIEmbedding struct {
	Client *openai.Client
}

func NewOpenAIEmbedding(client *openai.Client) *OpenAIEmbedding {
	return &OpenAIEmbedding{Client: client}
}

func (o *OpenAIEmbedding) GetEmbedding(text string) ([]float32, error) {
	req := openai.EmbeddingRequest{
		Input: []string{text},
		Model: openai.AdaEmbeddingV2,
	}

	resp, err := o.Client.CreateEmbeddings(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp.Data[0].Embedding, nil
}
