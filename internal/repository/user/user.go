package user

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

func NewRepo(store stores.Store) repository.UserRepo {
	var repo Repo
	repo.store = store
	return &repo
}

func (r Repo) Create(c context.Context, user models.User) (err error) {
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

func (r Repo) Get(c context.Context, user models.UserIdentity) (res models.User, err error) {
	bytes, err := r.store.Get(user.Name)
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

func (r Repo) Delete(c context.Context, user models.UserIdentity) (err error) {
	err = r.store.Del(user.Name)
	if errors.Is(err, appErr.ErrStoreRecNotFound) {
		err = appErr.ErrRepoRecNotFound
		return
	} else if err != nil {
		return
	}
	return
}
