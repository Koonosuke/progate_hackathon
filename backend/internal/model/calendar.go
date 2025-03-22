package model

import "time"

type CalendarDay struct {
	Date   time.Time `json:"date"`
	Retros []Retro   `json:"retros,omitempty"` // その日に複数のRetroがある場合、スライスで持たせる
}
