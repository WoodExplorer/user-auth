package services

import (
	"context"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/responses"
)

type User interface {
	Create(c context.Context, user requests.CreateUser) (err error)
	Get(c context.Context, user requests.GetUser) (res responses.GetUser, err error)
	Delete(c context.Context, user requests.DeleteUser) (err error)
}

type Role interface {
}

type UserRole interface {
}

type Authn interface {
}

type Authz interface {
}
