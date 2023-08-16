package repository

import (
	"context"
	"github.com/WoodExplorer/user-auth/internal/models"
)

type UserRepo interface {
	Create(c context.Context, user models.User) (err error)
	Get(c context.Context, user models.UserIdentity) (res models.User, err error)
	Delete(c context.Context, user models.UserIdentity) (err error)
	List(ctx context.Context) (res []models.User, err error)
}

type RoleRepo interface {
	Create(c context.Context, role models.Role) (err error)
	Get(c context.Context, role models.RoleIdentity) (res models.Role, err error)
	Delete(c context.Context, role models.RoleIdentity) (err error)
	List(ctx context.Context) (res []models.Role, err error)
}

type UserRoleRepo interface {
	Create(c context.Context, userRole models.UserRole) (err error)
}

type TokenBlacklistRepo interface {
	Create(c context.Context, token models.Token) (err error)
}
