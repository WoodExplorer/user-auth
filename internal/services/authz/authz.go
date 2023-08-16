package authz

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
	userRoleRepo  repository.UserRoleRepo
	blacklistRepo repository.TokenBlacklistRepo
}

func NewService(userRoleRepo repository.UserRoleRepo, blacklistRepo repository.TokenBlacklistRepo) services.Authz {
	var srv Service
	srv.userRoleRepo = userRoleRepo
	srv.blacklistRepo = blacklistRepo
	return &srv
}

func (s Service) CheckRole(c context.Context, r requests.CheckRole) (res responses.CheckRole, err error) {

	userInfo, err := pkg.ParseToken(r.Token, configs.GetJwtKey())
	if err != nil {
		err = errors.ErrAuthzInvalidToken
		return
	}

	existed, err := s.blacklistRepo.Exists(c, models.TokenIdentity{Value: r.Token})
	if err != nil {
		return
	}
	if existed {
		err = errors.ErrAuthzTokenInBlacklist
		return
	}

	existed, err = s.userRoleRepo.Exists(c, models.UserRoleIdentity{
		UserName: userInfo.Name,
		RoleName: r.RoleName,
	})
	if err != nil {
		return
	}

	res.Ok = existed
	return
}

func (s Service) AllRoles(c context.Context, r requests.AllRoles) (res responses.AllRoles, err error) {
	//TODO implement me
	panic("implement me")
}
