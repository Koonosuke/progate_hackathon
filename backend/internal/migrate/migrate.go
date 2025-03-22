package migrate

import (
	"database/sql"
	"log"

	"github.com/matthewTechCom/progate_hackathon/internal/config"
	"github.com/matthewTechCom/progate_hackathon/internal/infrastructure/db"
)

// Migrate データベースのマイグレーションを実行する
func Migrate(cfg *config.Config) {
	// db.goのInitDB関数を呼び出して、DB接続を取得
	dbConn := db.InitDB(cfg)

	// マイグレーション用のSQLを実行
	err := createBoardSummariesTable(dbConn)
	if err != nil {
		log.Fatalf("テーブル作成に失敗しました: %v", err)
	} else {
		log.Println("テーブル作成が完了しました")
	}
}

// createBoardSummariesTable テーブルを作成するSQL
func createBoardSummariesTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS board_summaries (
			id INT AUTO_INCREMENT PRIMARY KEY,
			summary TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	return err
}

func createStickyTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS sticky (
			id SERIAL PRIMARY KEY,
			board_id INT NOT NULL,
			miro_sticky_id TEXT NOT NULL,
			content TEXT NOT NULL,
			category TEXT CHECK (category IN ('改善点', '反省点')) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (board_id) REFERENCES board(id) ON DELETE CASCADE
		);
	`
	_, err := db.Exec(query)
	return err
}
