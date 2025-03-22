package model

import "time"

// ボード情報を表す
type Board struct {
	ID          int       `json:"id"`  // 参照されるやつ
	MiroBoardID string    `json:"miro_board_id"` // MiroのボードID
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// 付箋の詳しい情報を表す
type Sticky struct {
	ID           int       `json:"id"`
	BoardID      int       `json:"board_id"`       // BoardのIDを参照する（外部キー）
	MiroStickyID string    `json:"miro_sticky_id"` // Miroの付箋ID
	Content      string    `json:"content"`
	Category     string    `json:"category"`       // 改善 or 反省
	CreatedAt    time.Time `json:"created_at"`
	Embedding     []float32 
}