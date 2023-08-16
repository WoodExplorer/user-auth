package services

import (
	"context"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/responses"
)

type User interface {
	Create(c context.Context, user requests.CreateUser) (err error)
	Get(c context.Context, user requests.GetUser) (res responses.GetUser, err error)
	List(c context.Context) (res responses.ListUsers, err error)
	Delete(c context.Context, user requests.DeleteUser) (err error)
}

type Role interface {
	Create(c context.Context, user requests.CreateRole) (err error)
	Get(c context.Context, user requests.GetRole) (res responses.GetRole, err error)
	Delete(c context.Context, user requests.DeleteRole) (err error)
}

type UserRole interface {
}

type Authn interface {
}

type Authz interface {
}
