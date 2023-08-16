package user_role

import (
	"context"
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	"github.com/WoodExplorer/user-auth/internal/models"
	"github.com/WoodExplorer/user-auth/internal/repository"
	"github.com/WoodExplorer/user-auth/internal/stores"
	"github.com/pkg/errors"
)

const (
	keyPrefix = "tokenBlacklist:"
)

func prefix(key string) string {
	return keyPrefix + key
}

func getKey(t models.Token) string {
	return prefix(t.Value)
}

type Repo struct {
	store stores.Store
}

func NewRepo(store stores.Store) repository.TokenBlacklistRepo {
	var repo Repo
	repo.store = store
	return &repo
}

func (r Repo) Create(_ context.Context, token models.Token) (err error) {
	err = r.store.Set(getKey(token), nil)
	if err != nil {
		if errors.Is(err, appErr.ErrStoreRecAlreadyExists) {
			err = nil
		}
		return
	}
	return
}
