package errors

import "github.com/pkg/errors"

// store

var (
	ErrStoreRecAlreadyExists = errors.New("store: record already exists")
	ErrStoreRecNotFound      = errors.New("store: record not found")
	ErrStoreSubKeyNotFound   = errors.New("store: sub-key not found")
)

// repository

var (
	ErrRepoRecNotFound = errors.New("repo: record not found")
)

// service

var (
	ErrSvcUserAlreadyExisted = errors.New("user already existed")
	ErrSvcUserNotExisted     = errors.New("user does not existed")

	ErrSvcRoleAlreadyExisted = errors.New("role already existed")
	ErrSvcRoleNotExisted     = errors.New("role does not existed")
)

// authn
var (
	ErrAuthnFailed       = errors.New("authentication failed")
	ErrAuthnInvalidToken = errors.New("invalid token")
)

// authz
var (
	ErrAuthzInvalidToken     = errors.New("invalid token")
	ErrAuthzTokenInBlacklist = errors.New("token has been invalidated")
)
