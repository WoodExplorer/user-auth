package user_role

import (
	"context"
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/stores"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	keyPrefix = "userRole:"
)

func prefix(key string) string {
	return keyPrefix + key
}

func getKey(iVal interface{}) string {
	switch val := iVal.(type) {
	case models.UserRole:
		return prefix(val.UserName + ":" + val.RoleName)
	case models.UserRoleIdentity:
		return prefix(val.UserName + ":" + val.RoleName)
	default:
		log.Warn().Msgf("unknown supported: %+v", val)
	}
	return ""
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

func (r Repo) Exists(_ context.Context, userRole models.UserRoleIdentity) (ok bool, err error) {
	_, err = r.store.Get(getKey(userRole))
	if errors.Is(err, appErr.ErrStoreRecNotFound) {
		err = nil
		return
	} else if err != nil {
		return
	}
	ok = true
	return
}
