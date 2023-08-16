package memory

import (
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	"github.com/WoodExplorer/user-auth/internal/stores"
)

type Store struct {
	area map[string][]byte
}

func NewStore() stores.Store {
	var s Store
	s.area = make(map[string][]byte)
	return &s
}

func (s Store) Set(key string, data []byte) (err error) {
	s.area[key] = data
	return nil
}

func (s Store) SetNx(key string, data []byte) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) Get(key string) (res []byte, err error) {
	res, _ = s.area[key]
	return
}

func (s Store) GetE(key string) (res []byte, err error) {
	res, ok := s.area[key]
	if !ok {
		err = appErr.ErrStoreRecNotFound
		return
	}
	return
}

func (s Store) Del(key string) (err error) {
	delete(s.area, key)
	return nil
}

func (s Store) DelE(key string) (err error) {
	_, ok := s.area[key]
	if !ok {
		err = appErr.ErrStoreRecNotFound
		return
	}
	return s.Del(key)
}
