package pkg_test

import (
	"github.com/WoodExplorer/user-auth/internal/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	dummySignKey = "dummyKey"
)

var (
	dummyUser = pkg.UserInfo{Name: "jack"}
)

func TestTokenLogic(t *testing.T) {
	actual, err := pkg.CreateToken(dummyUser, dummySignKey)
	assert.Equal(t, err, nil)

	actualUserInfo, err := pkg.ParseToken(actual, dummySignKey)
	assert.Equal(t, err, nil)

	assert.Equal(t, actualUserInfo, dummyUser)
}
