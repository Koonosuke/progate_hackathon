package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matthewTechCom/progate_hackathon/internal/model"
)

type StickyRepositoryInterface interface {
	Save(stickies []*model.Sticky) ([]int, error)
}

type StickyRepository struct {
	DB *pgxpool.Pool
}

func NewStickyRepository(db *pgxpool.Pool) StickyRepositoryInterface {
	return &StickyRepository{DB: db}
}

// ヘルパー関数：[]float64 → '[...]' 形式の文字列に変換
func float64SliceToPGVector(vec []float64) string {
	strVals := make([]string, len(vec))
	for i, val := range vec {
		strVals[i] = fmt.Sprintf("%f", val)
	}
	return fmt.Sprintf("[%s]", strings.Join(strVals, ","))
}

func (r *StickyRepository) Save(stickies []*model.Sticky) ([]int, error) {
	var savedIDs []int

	for _, sticky := range stickies {
		query := `
			INSERT INTO sticky (board_id, miro_sticky_id, content, category, created_at, embedding)
			VALUES ($1, $2, $3, $4, DEFAULT, $5::vector)
			RETURNING id
		`

		// embeddingを []float64 に変換
		embedding := make([]float64, len(sticky.Embedding))
		for i, v := range sticky.Embedding {
			embedding[i] = float64(v)
		}
		embeddingStr := float64SliceToPGVector(embedding)

		var id int
		err := r.DB.QueryRow(context.Background(), query,
			sticky.BoardID,
			sticky.MiroStickyID,
			sticky.Content,
			sticky.Category,
			embeddingStr, // vector文字列形式で渡す
		).Scan(&id)

		if err != nil {
			return nil, err
		}
		savedIDs = append(savedIDs, id)
	}

	return savedIDs, nil
}
