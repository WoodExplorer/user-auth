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

type Repo struct {
	store stores.Store
}

func NewRepo(store stores.Store) repository.RoleRepo {
	var repo Repo
	repo.store = store
	return &repo
}

func (r Repo) Create(c context.Context, user models.Role) (err error) {
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	err = r.store.Set(user.Name, bytes)
	if err != nil {
		return
	}
	return
}

func (r Repo) Get(c context.Context, user models.RoleIdentity) (res models.Role, err error) {
	bytes, err := r.store.GetE(user.Name)
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

func (r Repo) Delete(c context.Context, user models.RoleIdentity) (err error) {
	err = r.store.DelE(user.Name)
	if errors.Is(err, appErr.ErrStoreRecNotFound) {
		err = appErr.ErrRepoRecNotFound
		return
	} else if err != nil {
		return
	}
	return
}
