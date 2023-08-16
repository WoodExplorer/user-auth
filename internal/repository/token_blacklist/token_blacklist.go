package token_blacklist

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
	keyPrefix = "tokenBlacklist:"
)

func prefix(key string) string {
	return keyPrefix + key
}

func getKey(iVal interface{}) string {
	switch val := iVal.(type) {
	case models.Token:
		return prefix(val.Value)
	case models.TokenIdentity:
		return prefix(val.Value)
	default:
		log.Warn().Msgf("unknown supported: %+v", val)
	}
	return ""
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

func (r Repo) Exists(_ context.Context, token models.TokenIdentity) (ok bool, err error) {
	_, err = r.store.Get(getKey(token))
	if errors.Is(err, appErr.ErrStoreRecNotFound) {
		err = nil
		return
	} else if err != nil {
		return
	}
	ok = true
	return
}
