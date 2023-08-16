package user_role

import (
	"context"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/services"
)

type Service struct {
	repo repository.UserRoleRepo
}

func NewService(repo repository.UserRoleRepo) services.UserRole {
	var srv Service
	srv.repo = repo
	return &srv
}

func (s *Service) Bind(c context.Context, req requests.BindUserRole) (err error) {

	// TODO: 存在性检查

	err = s.repo.Create(c, models.UserRole{UserName: req.UserName, RoleName: req.RoleName})
	if err != nil {
		return
	}
	return
}
