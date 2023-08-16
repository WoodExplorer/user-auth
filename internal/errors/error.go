package errors

import "github.com/pkg/errors"

var (
	ErrStoreRecAlreadyExists = errors.New("store: record already exists")
	ErrStoreRecNotFound      = errors.New("store: record not found")
)

var (
	ErrRepoRecNotFound = errors.New("repo: record not found")
)

var (
	ErrSvcUserExisted = errors.New("user already existed")
)

// authn
var (
	ErrAuthnFailed       = errors.New("authentication failed")
	ErrAuthnInvalidToken = errors.New("invalid token")
)
