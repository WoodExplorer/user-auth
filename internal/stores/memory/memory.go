package memory

import (
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	"github.com/WoodExplorer/user-auth/internal/stores"
	"github.com/pkg/errors"
)

const (
	opSet = "set"
	opGet = "get"
	opDel = "del"
)

type result struct {
	Err  error
	Data []byte
}

type command struct {
	Op   string
	Key  string
	Data []byte
	Ret  chan result
}

type Store struct {
	chCmd    chan command
	chClose  chan struct{}
	chExited chan struct{}
	area     map[string][]byte
}

func NewStore() stores.Store {
	var s Store
	s.chCmd = make(chan command, 1024)
	s.chClose = make(chan struct{})
	s.chExited = make(chan struct{})
	s.area = make(map[string][]byte)
	return &s
}

func (s *Store) Start() {
	go func() {
		for {
			select {
			case cmd := <-s.chCmd:
				cmd.Ret <- s.handleCmd(cmd)
			case <-s.chClose:
				s.chExited <- struct{}{}
			}
		}
	}()
}

func (s *Store) Stop() {
	close(s.chClose)
	<-s.chExited
}

func (s *Store) handleCmd(c command) (res result) {
	switch c.Op {
	case opSet:
		s.set(c.Key, c.Data)
	case opGet:
		res.Data, res.Err = s.get(c.Key)
	case opDel:
		res.Err = s.del(c.Key)
	default:
		res.Err = errors.Errorf("unknown operation: %s", c.Op)
	}
	return
}

func (s *Store) Set(key string, data []byte) (err error) {
	ret := make(chan result, 1)
	s.chCmd <- command{
		Op:   opSet,
		Key:  key,
		Data: data,
		Ret:  ret,
	}
	res := <-ret
	return res.Err
}

func (s *Store) Get(key string) ([]byte, error) {
	ret := make(chan result, 1)
	s.chCmd <- command{
		Op:  opGet,
		Key: key,
		Ret: ret,
	}
	res := <-ret
	return res.Data, res.Err
}

func (s *Store) Del(key string) (err error) {
	ret := make(chan result, 1)
	s.chCmd <- command{
		Op:  opDel,
		Key: key,
		Ret: ret,
	}
	res := <-ret
	return res.Err
}

func (s *Store) set(key string, data []byte) {
	s.area[key] = data
}

func (s *Store) get(key string) (res []byte, err error) {
	res, ok := s.area[key]
	if !ok {
		err = appErr.ErrStoreRecNotFound
		return
	}
	return
}

func (s *Store) del(key string) (err error) {
	_, ok := s.area[key]
	if !ok {
		err = appErr.ErrStoreRecNotFound
		return
	}
	delete(s.area, key)
	return
}
