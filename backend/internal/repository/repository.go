package repository

import (
	"database/sql"
	"github.com/matthewTechCom/progate_hackathon/internal/model"
)

// リポジトリのinterface
type BoardSummaryRepositoryInterface interface {
	Save(summary *model.BoardSummary) (int ,error)
}

type BoardSummaryRepository struct {
	DB *sql.DB
}

// コンストラクタを定義
func NewBoardSummaryRepository(db *sql.DB) BoardSummaryRepositoryInterface {
	return &BoardSummaryRepository{DB: db}
}

func (r BoardSummaryRepository) Save(summary *model.BoardSummary) (int, error) {
	// SQL実行
	result, err := r.DB.Exec(
		"INSERT INTO board_summaries (summary, created_at) VALUES (?, ?)",
		summary.Summary, summary.CreatedAt,
	)

	if err != nil {
		return 0, err
	}

	// 保存されたIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}