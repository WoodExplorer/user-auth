package pkg_test

import (
	"github.com/WoodExplorer/user-auth/internal/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPasswordHash(t *testing.T) {
	hash := pkg.GetPasswordHash("ua", "pwd")
	assert.Equal(t, hash, "c635757d6d1151f24705781f80f75d58")
}
