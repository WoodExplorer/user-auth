package user

import (
	"context"
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/pkg"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/responses"
	"github.com/WoodExplorer/user-auth/internal/services"
	"github.com/pkg/errors"
)

type Service struct {
	repo         repository.UserRepo
	userRoleRepo repository.UserRoleRepo
}

func NewService(repo repository.UserRepo, userRoleRepo repository.UserRoleRepo) services.User {
	var srv Service
	srv.repo = repo
	srv.userRoleRepo = userRoleRepo
	return &srv
}

func (s *Service) Create(c context.Context, req requests.CreateUser) (err error) {
	hash := pkg.GetPasswordHash(req.Name, req.Password)
	err = s.repo.Create(c, models.User{Name: req.Name, PasswordHash: hash})
	if err != nil {
		if errors.Is(err, appErr.ErrStoreRecAlreadyExists) {
			err = appErr.ErrSvcUserAlreadyExisted
		}
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

func (s *Service) List(c context.Context) (res responses.ListUsers, err error) {
	users, err := s.repo.List(c)
	if err != nil {
		return
	}

	for _, u := range users {
		res.Items = append(res.Items, responses.User{Name: u.Name})
	}
	return
}

func (s *Service) Delete(c context.Context, req requests.DeleteUser) (err error) {
	err = s.repo.Delete(c, models.UserIdentity{Name: req.Name})
	if err != nil {
		if errors.Is(err, appErr.ErrRepoRecNotFound) {
			err = appErr.ErrSvcUserNotExisted
		}
		return
	}

	err = s.userRoleRepo.DeleteByUser(c, models.UserIdentity{Name: req.Name})
	if err != nil {
		return
	}

	return
}
