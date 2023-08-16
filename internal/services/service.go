package services

import (
	"context"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/responses"
)

type User interface {
	Create(c context.Context, r requests.CreateUser) (err error)
	Get(c context.Context, r requests.GetUser) (res responses.GetUser, err error)
	List(c context.Context) (res responses.ListUsers, err error)
	Delete(c context.Context, r requests.DeleteUser) (err error)
}

type Role interface {
	Create(c context.Context, r requests.CreateRole) (err error)
	Get(c context.Context, r requests.GetRole) (res responses.GetRole, err error)
	List(c context.Context) (res responses.ListRoles, err error)
	Delete(c context.Context, r requests.DeleteRole) (err error)
}

type UserRole interface {
	Bind(c context.Context, r requests.BindUserRole) (err error)
}

type Authn interface {
	Authenticate(c context.Context, r requests.Authenticate) (res responses.Authenticate, err error)
	Invalidate(c context.Context, r requests.Invalidate) (err error)
}

type Authz interface {
	CheckRole(c context.Context, r requests.CheckRole) (res responses.CheckRole, err error)
	GetUserRoles(c context.Context, r requests.UserRoles) (res responses.UserRoles, err error)
}
