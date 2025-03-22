package model

import "time"

type Retro struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"` // レトロを登録したuserのid
	Date      time.Time `json:"date"`
	MiroURL   string    `json:"miro_url"` // MiroのURL
	Summary   string    `json:"summary"`  // ChatGPTにより生成された要約
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
