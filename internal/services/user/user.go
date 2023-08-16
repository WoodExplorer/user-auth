package user

import (
	"context"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/responses"
	"github.com/WoodExplorer/user-auth/internal/services"
)

type Service struct {
	repo repository.UserRepo
}

func NewService(repo repository.UserRepo) services.User {
	var srv Service
	srv.repo = repo
	return &srv
}

func (s *Service) Create(c context.Context, req requests.CreateUser) (err error) {
	err = s.repo.Create(c, models.User{Name: req.Name, PasswordHash: req.Password})
	if err != nil {
		return
	}
	return
}

func (s *Service) Get(c context.Context, req requests.GetUser) (res responses.GetUser, err error) {
	model, err := s.repo.Get(c, models.UserIdentity{Name: req.Name})
	if err != nil {
		return
	}

	res.Name = model.Name
	return
}

func (s *Service) Delete(c context.Context, req requests.DeleteUser) (err error) {
	err = s.repo.Delete(c, models.UserIdentity{Name: req.Name})
	if err != nil {
		return
	}

	return
}
