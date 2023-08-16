package authn_test

import (
	"context"
	role_repo "github.com/WoodExplorer/user-auth/internal/repository/role"
	"github.com/WoodExplorer/user-auth/internal/repository/token_blacklist"
	user_repo "github.com/WoodExplorer/user-auth/internal/repository/user"
	user_role_repo "github.com/WoodExplorer/user-auth/internal/repository/user_role"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/services/authn"
	"github.com/WoodExplorer/user-auth/internal/services/role"
	"github.com/WoodExplorer/user-auth/internal/services/user"
	"github.com/WoodExplorer/user-auth/internal/services/user_role"
	"github.com/WoodExplorer/user-auth/internal/stores/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthn(t *testing.T) {
	var err error
	c := context.Background()

	store := memory.NewStore()
	store.Start()
	defer store.Stop()

	ur := user_repo.NewRepo(store)
	us := user.NewService(ur)

	rr := role_repo.NewRepo(store)
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

	urr := user_role_repo.NewRepo(store)
	urs := user_role.NewService(urr)
	err = urs.Bind(c, requests.BindUserRole{
		UserName: "ua",
		RoleName: "ra",
	})
	assert.Equal(t, err, nil)

	tbr := token_blacklist.NewRepo(store)
	ans := authn.NewService(ur, tbr)

	token, err := ans.Authenticate(c, requests.Authenticate{
		Name:     "ua",
		Password: "pwd",
	})
	assert.Equal(t, err, nil)
	assert.True(t, len(token.Token) > 0)

	err = ans.Invalidate(c, requests.Invalidate{
		Token: token.Token,
	})
	assert.Equal(t, err, nil)
}
