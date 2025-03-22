package db

import (
	"database/sql"
	"fmt"
	"log"
	
	"github.com/matthewTechCom/progate_hackathon/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLへの接続を初期化して*sql.DBを返す
func InitDB(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
	cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("MySQL接続に失敗しました: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("MySQLへのPingに失敗しました: %v", err)
	}

	// // オプション：接続プールの設定
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(5)

	log.Println("MySQL接続に成功しました")
	return db

}