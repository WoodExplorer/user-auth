package user_role

import (
	"context"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/stores"
)

const (
	keyPrefix = "userRole:"
)

func prefix(key string) string {
	return keyPrefix + key
}

func getKey(userRole models.UserRole) string {
	return prefix(userRole.UserName + ":" + userRole.RoleName)
}

type Repo struct {
	store stores.Store
}

func NewRepo(store stores.Store) repository.UserRoleRepo {
	var repo Repo
	repo.store = store
	return &repo
}

func (r Repo) Create(_ context.Context, userRole models.UserRole) (err error) {
	err = r.store.Set(getKey(userRole), nil)
	if err != nil {
		return
	}
	return
}
