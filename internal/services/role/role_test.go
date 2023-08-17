package role_test

import (
	"context"
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	role_repo "github.com/WoodExplorer/user-auth/internal/repository/role"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/services/role"
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

	rr := role_repo.NewRepo(store)
	svc := role.NewService(rr)

	err = svc.Create(c, requests.CreateRole{
		Name: "ra",
	})
	assert.Equal(t, err, nil)

	err = svc.Create(c, requests.CreateRole{
		Name: "ra",
	})
	assert.True(t, errors.Is(err, appErr.ErrSvcRoleAlreadyExisted))
}

func TestDelete(t *testing.T) {
	var err error
	c := context.Background()

	store := memory.NewStore()
	store.Start()
	defer store.Stop()

	rr := role_repo.NewRepo(store)
	svc := role.NewService(rr)

	err = svc.Create(c, requests.CreateRole{
		Name: "ra",
	})
	assert.Equal(t, err, nil)

	err = svc.Delete(c, requests.DeleteRole{
		Name: "ra",
	})
	assert.Equal(t, err, nil)

	err = svc.Delete(c, requests.DeleteRole{
		Name: "ua",
	})
	assert.True(t, errors.Is(err, appErr.ErrSvcRoleNotExisted))
}
