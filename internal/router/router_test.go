package router_test

import (
	"bytes"
	"encoding/json"
	role_repo "github.com/WoodExplorer/user-auth/internal/repository/role"
	token_blacklist_repo "github.com/WoodExplorer/user-auth/internal/repository/token_blacklist"
	user_repo "github.com/WoodExplorer/user-auth/internal/repository/user"
	user_role_repo "github.com/WoodExplorer/user-auth/internal/repository/user_role"
	"github.com/WoodExplorer/user-auth/internal/responses"
	"github.com/WoodExplorer/user-auth/internal/router"
	"github.com/WoodExplorer/user-auth/internal/services/authn"
	"github.com/WoodExplorer/user-auth/internal/services/authz"
	"github.com/WoodExplorer/user-auth/internal/services/role"
	"github.com/WoodExplorer/user-auth/internal/services/user"
	"github.com/WoodExplorer/user-auth/internal/services/user_role"
	"github.com/WoodExplorer/user-auth/internal/stores/memory"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
)

func TestRouter(t *testing.T) {

	store := memory.NewStore()
	store.Start()

	rr := role_repo.NewRepo(store)
	ur := user_repo.NewRepo(store)
	urr := user_role_repo.NewRepo(store)
	tbr := token_blacklist_repo.NewRepo(store)

	roleSvc := role.NewService(rr)
	userSvc := user.NewService(ur, urr)
	userRoleSvc := user_role.NewService(urr)
	authnSvc := authn.NewService(ur, tbr)
	authzSvc := authz.NewService(urr, tbr)

	eng, _ := router.InitRouter(roleSvc, userSvc, userRoleSvc, authnSvc, authzSvc)

	res := doReq(t, eng, http.MethodPost, "/api/v1/users/", buildBytes(map[string]interface{}{
		"name":     "ua",
		"password": "pwd",
	}))
	assert.Equal(t, res.Code, 0)

	res = doReq(t, eng, http.MethodPost, "/api/v1/roles/", buildBytes(map[string]interface{}{
		"name": "ra",
	}))
	assert.Equal(t, res.Code, 0)

	res = doReq(t, eng, http.MethodPost, "/api/v1/user-roles/", buildBytes(map[string]interface{}{
		"userName": "ua",
		"roleName": "ra",
	}))
	assert.Equal(t, res.Code, 0)

	res = doReq(t, eng, http.MethodPost, "/api/v1/authn/tokens", buildBytes(map[string]interface{}{
		"name":     "ua",
		"password": "pwd",
	}))
	assert.Equal(t, res.Code, 0)
	var token responses.Authenticate
	err := mapstructure.Decode(res.Data, &token)
	assert.Equal(t, err, nil)
	assert.True(t, len(token.Token) > 0)

	res = doReq(t, eng, http.MethodGet, "/api/v1/authz/check-role?token="+token.Token+"&roleName="+"ra", nil)
	assert.Equal(t, res.Code, 0)
	var check responses.CheckRole
	err = mapstructure.Decode(res.Data, &check)
	assert.Equal(t, err, nil)
	assert.True(t, check.Ok)

	res = doReq(t, eng, http.MethodGet, "/api/v1/authz/user-roles?token="+token.Token, nil)
	assert.Equal(t, res.Code, 0)
	var userRoles responses.UserRoles
	err = mapstructure.Decode(res.Data, &userRoles)
	assert.Equal(t, err, nil)
	sort.Slice(userRoles.Roles, func(i, j int) bool {
		return userRoles.Roles[i].Name < userRoles.Roles[j].Name
	})
	assert.Equal(t, responses.UserRoles{Roles: []responses.Role{
		{"ra"},
	}}, userRoles)
}

func buildBytes(data map[string]interface{}) io.Reader {
	bs, _ := json.Marshal(data)
	return bytes.NewReader(bs)
}

func doReq(t *testing.T, eng *gin.Engine, method, url string, body io.Reader) (wrapper responses.Wrapper) {

	req := httptest.NewRequest(method, url, body)

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	w := httptest.NewRecorder()
	eng.ServeHTTP(w, req)
	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	return
}
