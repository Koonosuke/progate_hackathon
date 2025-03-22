package repository

import (
	"github.com/matthewTechCom/progate_hackathon/internal/model"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	FindByGoogleID(googleID string) (*model.User, error)
	CreateUser(user *model.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{db: db}
}

func (ar *authRepository) FindByGoogleID(googleID string) (*model.User, error) {
	var user model.User
	if err := ar.db.Where("google_id = ?", googleID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar *authRepository) CreateUser(user *model.User) error {
	return ar.db.Create(user).Error
}
