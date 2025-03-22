package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matthewTechCom/progate_hackathon/internal/model"
)

type BoardRepositoryInterface interface {
	Save(board *model.Board) (int, error)
	GetByMiroID(miroBoardID string) (*model.Board, error)
}

type BoardRepository struct {
	DB *pgxpool.Pool
}

func NewBoardRepository(db *pgxpool.Pool) BoardRepositoryInterface {
	return &BoardRepository{DB: db}
}

func (r *BoardRepository) Save(board *model.Board) (int, error) {
	query := `
		INSERT INTO board (miro_board_id, title, description, created_at)
		VALUES ($1, $2, $3, DEFAULT)
		RETURNING id
	`
	var id int
	err := r.DB.QueryRow(context.Background(), query, board.MiroBoardID, board.Title, board.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BoardRepository) GetByMiroID(miroBoardID string) (*model.Board, error) {
	query := `
		SELECT id, miro_board_id, title, description, created_at
		FROM board
		WHERE miro_board_id = $1
	`
	row := r.DB.QueryRow(context.Background(), query, miroBoardID)
	var board model.Board
	err := row.Scan(&board.ID, &board.MiroBoardID, &board.Title, &board.Description, &board.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &board, nil
}
