package authz_test

import (
	"context"
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	role_repo "github.com/WoodExplorer/user-auth/internal/repository/role"
	"github.com/WoodExplorer/user-auth/internal/repository/token_blacklist"
	user_repo "github.com/WoodExplorer/user-auth/internal/repository/user"
	user_role_repo "github.com/WoodExplorer/user-auth/internal/repository/user_role"
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/WoodExplorer/user-auth/internal/responses"
	"github.com/WoodExplorer/user-auth/internal/services/authn"
	"github.com/WoodExplorer/user-auth/internal/services/authz"
	"github.com/WoodExplorer/user-auth/internal/services/role"
	"github.com/WoodExplorer/user-auth/internal/services/user"
	"github.com/WoodExplorer/user-auth/internal/services/user_role"
	"github.com/WoodExplorer/user-auth/internal/stores/memory"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestAuthz(t *testing.T) {
	var err error
	c := context.Background()

	store := memory.NewStore()
	store.Start()
	defer store.Stop()

	ur := user_repo.NewRepo(store)
	us := user.NewService(ur)

	rr := role_repo.NewRepo(store)
	rs := role.NewService(rr)

	urr := user_role_repo.NewRepo(store)
	urs := user_role.NewService(urr)

	tbr := token_blacklist.NewRepo(store)
	ans := authn.NewService(ur, tbr)

	azs := authz.NewService(urr, tbr)

	err = us.Create(c, requests.CreateUser{
		Name:     "ua",
		Password: "pwd",
	})
	assert.Equal(t, err, nil)

	err = rs.Create(c, requests.CreateRole{
		Name: "ra",
	})
	assert.Equal(t, err, nil)

	err = rs.Create(c, requests.CreateRole{
		Name: "rb",
	})
	assert.Equal(t, err, nil)

	err = urs.Bind(c, requests.BindUserRole{
		UserName: "ua",
		RoleName: "ra",
	})
	assert.Equal(t, err, nil)

	err = urs.Bind(c, requests.BindUserRole{
		UserName: "ua",
		RoleName: "rb",
	})
	assert.Equal(t, err, nil)

	token, err := ans.Authenticate(c, requests.Authenticate{
		Name:     "ua",
		Password: "pwd",
	})
	assert.Equal(t, err, nil)
	assert.True(t, len(token.Token) > 0)

	check, err := azs.CheckRole(c, requests.CheckRole{
		Token:    token.Token,
		RoleName: "ra",
	})
	assert.Equal(t, err, nil)
	assert.True(t, check.Ok)

	check, err = azs.CheckRole(c, requests.CheckRole{
		Token:    token.Token,
		RoleName: "rc",
	})
	assert.Equal(t, err, nil)
	assert.False(t, check.Ok)

	userRoles, err := azs.GetUserRoles(c, requests.UserRoles{
		Token: token.Token,
	})
	sort.Slice(userRoles.Roles, func(i, j int) bool {
		return userRoles.Roles[i].Name < userRoles.Roles[j].Name
	})
	assert.Equal(t, err, nil)
	assert.Equal(t, responses.UserRoles{Roles: []responses.Role{
		{"ra"}, {"rb"},
	}}, userRoles)

	err = ans.Invalidate(c, requests.Invalidate{
		Token: token.Token,
	})
	assert.Equal(t, err, nil)

	check, err = azs.CheckRole(c, requests.CheckRole{
		Token:    token.Token,
		RoleName: "ra",
	})
	assert.True(t, errors.Is(err, appErr.ErrAuthzTokenInBlacklist))
	assert.False(t, check.Ok)
}
