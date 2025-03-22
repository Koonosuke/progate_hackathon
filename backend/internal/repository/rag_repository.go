package repository

import (
	"context"
"fmt"

 "github.com/pgvector/pgvector-go"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RAGRepository struct {
	DB *pgxpool.Pool
}

// Embedding を更新
// ✅ 修正後：[]float64 をそのまま渡す
func (r *RAGRepository) UpdateEmbedding(id int, embedding []float64) error {
	// pgvector-go は float32 のスライスを受け取るので変換
	vector := pgvector.NewVector(convertToFloat32(embedding))

	_, err := r.DB.Exec(context.Background(),
		`UPDATE stickies SET embedding = $1 WHERE id = $2`, vector, id)
	return err
}

// ヘルパー関数：[]float64 → []float32 に変換
func convertToFloat32(arr []float64) []float32 {
	f32 := make([]float32, len(arr))
	for i, v := range arr {
		f32[i] = float32(v)
	}
	return f32
}



// ベクトル検索を実行
func (repo *RAGRepository) SearchStickies(embedding []float64, category string) ([]map[string]string, error) {
	// 🔧 embedding を pgvector.Vector に変換（float64 → float32）
	vector := pgvector.NewVector(convertToFloat32(embedding))

	sqlQuery := `
		SELECT content, category, color, created_at
		FROM stickies
		WHERE embedding IS NOT NULL
	`
	args := []interface{}{vector}
	argIndex := 2

	// category フィルタがある場合
	if category != "" {
		sqlQuery += fmt.Sprintf(" AND category = $%d", argIndex)
		args = append(args, category)
		argIndex++
	}

	sqlQuery += " ORDER BY embedding <-> $1 LIMIT 5;"

	rows, err := repo.DB.Query(context.Background(), sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]string
	for rows.Next() {
		var content, category, color, createdAt string
		rows.Scan(&content, &category, &color, &createdAt)
		results = append(results, map[string]string{
			"content":    content,
			"category":   category,
			"color":      color,
			"created_at": createdAt,
		})
	}

	return results, nil
}
