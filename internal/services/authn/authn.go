package authn

import (
	"context"
	"github.com/WoodExplorer/user-auth/internal/configs"
	"github.com/WoodExplorer/user-auth/internal/errors"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/pkg"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/responses"
	"github.com/WoodExplorer/user-auth/internal/services"
)

type Service struct {
	repo repository.UserRepo
}

func NewService(repo repository.UserRepo) services.Authn {
	var srv Service
	srv.repo = repo
	return &srv
}

func (s *Service) Authenticate(c context.Context, req requests.Authenticate) (res responses.Authenticate, err error) {
	user, err := s.repo.Get(c, models.UserIdentity{Name: req.Name})
	if err != nil {
		err = errors.ErrAuthnFailed
		return
	}

	actualHash := pkg.GetPasswordHash(req.Name, req.Password)
	if user.PasswordHash != actualHash {
		err = errors.ErrAuthnFailed
		return
	}

	token, err := pkg.CreateToken(pkg.UserInfo{Name: req.Name}, configs.GetJwtKey())
	if err != nil {
		return
	}

	res.Token = token
	return
}

func (s *Service) Invalidate(c context.Context, r requests.Invalidate) (err error) {
	//TODO implement me
	panic("implement me")
}
