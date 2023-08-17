package role

import (
	"context"
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/responses"
	"github.com/WoodExplorer/user-auth/internal/services"
	"github.com/pkg/errors"
)

type Service struct {
	repo repository.RoleRepo
}

func NewService(repo repository.RoleRepo) services.Role {
	var srv Service
	srv.repo = repo
	return &srv
}

func (s *Service) Create(c context.Context, req requests.CreateRole) (err error) {
	err = s.repo.Create(c, models.Role{Name: req.Name})
	if err != nil {
		if errors.Is(err, appErr.ErrStoreRecAlreadyExists) {
			err = appErr.ErrSvcRoleAlreadyExisted
		}
		return
	}
	return
}

func (s *Service) Get(c context.Context, req requests.GetRole) (res responses.GetRole, err error) {
	model, err := s.repo.Get(c, models.RoleIdentity{Name: req.Name})
	if err != nil {
		return
	}

	res.Name = model.Name
	return
}

func (s *Service) List(c context.Context) (res responses.ListRoles, err error) {
	users, err := s.repo.List(c)
	if err != nil {
		return
	}

	for _, u := range users {
		res.Items = append(res.Items, responses.Role{Name: u.Name})
	}
	return
}

func (s *Service) Delete(c context.Context, req requests.DeleteRole) (err error) {
	err = s.repo.Delete(c, models.RoleIdentity{Name: req.Name})
	if err != nil {
		if errors.Is(err, appErr.ErrRepoRecNotFound) {
			err = appErr.ErrSvcRoleNotExisted
		}
		return
	}
	return
}
