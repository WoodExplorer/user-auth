package user_role

import (
	"context"
	"encoding/json"
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
	case models.User:
		return prefix(val.Name)
	case models.UserIdentity:
		return prefix(val.Name)
	case models.UserRole:
		return prefix(val.UserName)
	case models.UserRoleIdentity:
		return prefix(val.UserName)
	default:
		log.Warn().Msgf("unknown supported: %+v", val)
	}
	return ""
}

func getSubKey(iVal interface{}) string {
	switch val := iVal.(type) {
	case models.UserRole:
		return val.RoleName
	case models.UserRoleIdentity:
		return val.RoleName
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
	bytes, err := json.Marshal(userRole)
	if err != nil {
		return
	}

	err = r.store.HSet(getKey(userRole), getSubKey(userRole), bytes)
	if err != nil {
		return
	}

	return
}

func (r Repo) Exists(_ context.Context, userRole models.UserRoleIdentity) (ok bool, err error) {
	_, err = r.store.HGet(getKey(userRole), getSubKey(userRole))
	if errors.Is(err, appErr.ErrStoreRecNotFound) || errors.Is(err, appErr.ErrStoreSubKeyNotFound) {
		err = nil
		return
	} else if err != nil {
		return
	}

	ok = true
	return
}

func (r Repo) GetUserRoles(_ context.Context, user models.UserIdentity) (res []models.UserRole, err error) {
	item, err := r.store.HGetAll(getKey(user))
	if errors.Is(err, appErr.ErrStoreRecNotFound) {
		err = nil
		return
	} else if err != nil {
		return
	}

	for _, val := range item {
		var buf models.UserRole
		err = json.Unmarshal(val, &buf)
		if err != nil {
			return
		}
		res = append(res, buf)
	}
	return
}

func (r Repo) DeleteByUser(_ context.Context, user models.UserIdentity) (err error) {
	err = r.store.HDelAll(getKey(user))
	if errors.Is(err, appErr.ErrStoreRecNotFound) {
		err = nil // TODO: inconsistent behavior regarding to delete
		return
	} else if err != nil {
		return
	}

	return
}
