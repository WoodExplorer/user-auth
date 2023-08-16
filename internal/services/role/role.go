package role

import (
	"context"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/responses"
	"github.com/WoodExplorer/user-auth/internal/services"
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

func (s *Service) Delete(c context.Context, req requests.DeleteRole) (err error) {
	err = s.repo.Delete(c, models.RoleIdentity{Name: req.Name})
	if err != nil {
		return
	}

	return
}
