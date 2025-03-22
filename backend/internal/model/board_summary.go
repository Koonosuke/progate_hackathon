package model

import "time"

type BoardSummary struct {
	ID int `json:"id"`
	Summary string `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
}