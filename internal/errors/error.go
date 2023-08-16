package errors

import "github.com/pkg/errors"

var (
	ErrStoreRecNotFound = errors.New("store: record not found")
)

var (
	ErrRepoRecNotFound = errors.New("repo: record not found")
)

var (
	ErrUserExisted = errors.New("user already existed")
)
