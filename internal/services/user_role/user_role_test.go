package user_role_test

import (
	"context"
	role_repo "github.com/WoodExplorer/user-auth/internal/repository/role"
	user_repo "github.com/WoodExplorer/user-auth/internal/repository/user"
	user_role_repo "github.com/WoodExplorer/user-auth/internal/repository/user_role"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/services/role"
	"github.com/WoodExplorer/user-auth/internal/services/user"
	"github.com/WoodExplorer/user-auth/internal/services/user_role"
	"github.com/WoodExplorer/user-auth/internal/stores/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBind(t *testing.T) {
	var err error
	c := context.Background()

	store := memory.NewStore()
	store.Start()
	defer store.Stop()

	ur := user_repo.NewRepo(store)
	rr := role_repo.NewRepo(store)
	urr := user_role_repo.NewRepo(store)

	us := user.NewService(ur, urr)
	rs := role.NewService(rr)

	err = us.Create(c, requests.CreateUser{
		Name:     "ua",
		Password: "pwd",
	})
	assert.Equal(t, err, nil)

	err = rs.Create(c, requests.CreateRole{
		Name: "ra",
	})
	assert.Equal(t, err, nil)

	urs := user_role.NewService(urr)
	err = urs.Bind(c, requests.BindUserRole{
		UserName: "ua",
		RoleName: "ra",
	})
	assert.Equal(t, err, nil)
}
