package repository

import (
    "context"
    "fmt"
    "time"

    "github.com/pgvector/pgvector-go"
    "github.com/jackc/pgx/v5/pgxpool"
)

type RAGRepository struct {
    DB *pgxpool.Pool
}

// Embedding を更新
func (r *RAGRepository) UpdateEmbedding(id int, embedding []float64) error {
    // pgvector-go は float32 のスライスを受け取るので変換
    vector := pgvector.NewVector(convertToFloat32(embedding))

    _, err := r.DB.Exec(context.Background(),
        `UPDATE sticky SET embedding = $1 WHERE id = $2`, vector, id)
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

// `content` で RAG を実行し、カテゴリでフィルタ＆ created_at で並び替え
func (repo *RAGRepository) SearchStickies(embedding []float64, category string) ([]map[string]string, error) {
    // 🔧 embedding を pgvector.Vector に変換（float64 → float32）
    vector := pgvector.NewVector(convertToFloat32(embedding))

    sqlQuery := `
        SELECT content, category, created_at
        FROM sticky
        WHERE embedding IS NOT NULL
    `

    args := []interface{}{vector}
    argIndex := 2

    // category フィルタを適用
    if category != "" {
        sqlQuery += fmt.Sprintf(" AND category = $%d", argIndex)
        args = append(args, category)
        argIndex++
    }

    // 🔥 ORDER BY を修正して、ベクトル検索結果を優先しつつ、created_at の新しい順に並べる
    sqlQuery += " ORDER BY embedding <-> $1, created_at DESC LIMIT 5;"

    rows, err := repo.DB.Query(context.Background(), sqlQuery, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []map[string]string
    for rows.Next() {
        var content, category string
        var createdAt time.Time // `created_at` は `time.Time` 型として取得

        if err := rows.Scan(&content, &category,  &createdAt); err != nil {
            return nil, err
        }

        results = append(results, map[string]string{
            "content":    content,
            "category":   category,
           
            "created_at": createdAt.Format("2006-01-02 15:04:05"), // `time.Time` を `string` に変換
        })
    }

    return results, nil
}
