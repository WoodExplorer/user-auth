package user_test

import (
	"context"
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	user_repo "github.com/WoodExplorer/user-auth/internal/repository/user"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/services/user"
	"github.com/WoodExplorer/user-auth/internal/stores/memory"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	var err error
	c := context.Background()

	store := memory.NewStore()
	store.Start()
	defer store.Stop()

	ur := user_repo.NewRepo(store)
	svc := user.NewService(ur)

	err = svc.Create(c, requests.CreateUser{
		Name:     "ua",
		Password: "pwd",
	})
	assert.Equal(t, err, nil)

	err = svc.Create(c, requests.CreateUser{
		Name:     "ua",
		Password: "pwd",
	})
	assert.True(t, errors.Is(err, appErr.ErrSvcUserAlreadyExisted))
}

func TestDelete(t *testing.T) {
	var err error
	c := context.Background()

	store := memory.NewStore()
	store.Start()
	defer store.Stop()

	ur := user_repo.NewRepo(store)
	svc := user.NewService(ur)

	err = svc.Create(c, requests.CreateUser{
		Name:     "ua",
		Password: "pwd",
	})
	assert.Equal(t, err, nil)

	err = svc.Delete(c, requests.DeleteUser{
		Name: "ua",
	})
	assert.Equal(t, err, nil)

	err = svc.Delete(c, requests.DeleteUser{
		Name: "ua",
	})
	assert.True(t, errors.Is(err, appErr.ErrSvcUserNotExisted))
}
