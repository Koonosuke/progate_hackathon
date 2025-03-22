package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/matthewTechCom/progate_hackathon/internal/model"
	"github.com/matthewTechCom/progate_hackathon/internal/repository"
	"github.com/matthewTechCom/progate_hackathon/util"
	"gorm.io/gorm"
)

type IAuthUsecase interface {
	HandleGoogleCallback(ctx context.Context, code string) (*model.User, error)
}

type authUsecase struct {
	ar repository.IAuthRepository
}

func NewAuthUsecase(ar repository.IAuthRepository) IAuthUsecase {
	return &authUsecase{
		ar: ar,
	}
}

func (au *authUsecase) HandleGoogleCallback(ctx context.Context, code string) (*model.User, error) {
	token, err := util.GoogleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := util.GoogleOauthConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
	if err := json.Unmarshal(data, &userInfo); err != nil {
		return nil, err
	}

	user, err := au.ar.FindByGoogleID(userInfo.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user = &model.User{
				GoogleID: userInfo.ID,
				Name:     userInfo.Name,
				Email:    userInfo.Email,
				Picture:  userInfo.Picture,
			}
			if err := au.ar.CreateUser(user); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return user, nil
}
