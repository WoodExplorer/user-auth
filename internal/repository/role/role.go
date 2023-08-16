package role

import (
	"context"
	"encoding/json"
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/stores"
	"github.com/pkg/errors"
)

const (
	keyPrefix = "role-"
)

func prefix(key string) string {
	return keyPrefix + key
}

type Repo struct {
	store stores.Store
}

func NewRepo(store stores.Store) repository.RoleRepo {
	var repo Repo
	repo.store = store
	return &repo
}

func (r Repo) Create(_ context.Context, user models.Role) (err error) {
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	err = r.store.Set(prefix(user.Name), bytes)
	if err != nil {
		return
	}
	return
}

func (r Repo) Get(_ context.Context, user models.RoleIdentity) (res models.Role, err error) {
	bytes, err := r.store.Get(prefix(user.Name))
	if errors.Is(err, appErr.ErrStoreRecNotFound) {
		err = appErr.ErrRepoRecNotFound
		return
	} else if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return
	}
	return
}

func (r Repo) List(_ context.Context) (res []models.Role, err error) {
	bytesSlice, err := r.store.Keys(keyPrefix)
	if err != nil {
		return
	}
	for _, bytes := range bytesSlice {
		var buf models.Role
		err = json.Unmarshal(bytes, &buf)
		if err != nil {
			return
		}
		res = append(res, buf)
	}
	return
}

func (r Repo) Delete(_ context.Context, user models.RoleIdentity) (err error) {
	err = r.store.Del(prefix(user.Name))
	if errors.Is(err, appErr.ErrStoreRecNotFound) {
		err = appErr.ErrRepoRecNotFound
		return
	} else if err != nil {
		return
	}
	return
}
