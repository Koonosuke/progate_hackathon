package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	GoogleID string `gorm:"uniqueIndex"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Picture  string
	GroupIDs []int `json:"group_ids,omitempty"`
}
