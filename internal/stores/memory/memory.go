package memory

import (
	"encoding/json"
	appErr "github.com/WoodExplorer/user-auth/internal/errors"
	"github.com/WoodExplorer/user-auth/internal/stores"
	"github.com/pkg/errors"
	"strings"
)

type Store struct {
	chCmd    chan Command
	chClose  chan struct{}
	chExited chan struct{}
	area     map[string]interface{}
}

func NewStore() stores.Store {
	var s Store
	s.chCmd = make(chan Command, 1024)
	s.chClose = make(chan struct{})
	s.chExited = make(chan struct{})
	s.area = make(map[string]interface{})
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

func (s *Store) handleCmd(c Command) (res Result) {
	switch c.Op {
	case opSet:
		res.Err = s.set(c.Key, c.Data)
	case opGet:
		res.Data, res.Err = s.get(c.Key)
	case opDel:
		res.Err = s.del(c.Key)
	case opKeys:
		res.Data, res.Err = s.keys(c.Key)
	case opHSet:
		res.Err = s.hset(c.Key, c.SubKey, c.Data)
	case opHGet:
		res.Data, res.Err = s.hget(c.Key, c.SubKey)
	case opHGetAll:
		res.Data, res.Err = s.hgetall(c.Key)
	case opHDelAll:
		res.Err = s.hdelall(c.Key)
	case opBatch:
		res.Err = s.batch(c.Data)
	default:
		res.Err = errors.Errorf("unknown operation: %s", c.Op)
	}
	return
}

func (s *Store) BeginTx() stores.Store {
	tx := NewTx(s)
	return tx
}

func (s *Store) CommitTx() (err error) {
	// note: should not be called
	return
}

func (s *Store) Set(key string, data []byte) (err error) {
	ret := make(chan Result, 1)
	s.chCmd <- Command{
		Op:   opSet,
		Key:  key,
		Data: data,
		Ret:  ret,
	}
	res := <-ret
	return res.Err
}

func (s *Store) Get(key string) ([]byte, error) {
	ret := make(chan Result, 1)
	s.chCmd <- Command{
		Op:  opGet,
		Key: key,
		Ret: ret,
	}
	res := <-ret
	return res.Data.([]byte), res.Err
}

func (s *Store) Del(key string) (err error) {
	ret := make(chan Result, 1)
	s.chCmd <- Command{
		Op:  opDel,
		Key: key,
		Ret: ret,
	}
	res := <-ret
	return res.Err
}

func (s *Store) Keys(keyPrefix string) (data [][]byte, err error) {
	ret := make(chan Result, 1)
	s.chCmd <- Command{
		Op:  opKeys,
		Key: keyPrefix,
		Ret: ret,
	}
	res := <-ret
	return res.Data.([][]byte), res.Err
}

func (s *Store) HSet(key string, subKey string, data []byte) (err error) {
	ret := make(chan Result, 1)
	s.chCmd <- Command{
		Op:     opHSet,
		Key:    key,
		SubKey: subKey,
		Data:   data,
		Ret:    ret,
	}
	res := <-ret
	return res.Err
}

func (s *Store) HGet(key string, subKey string) (data []byte, err error) {
	ret := make(chan Result, 1)
	s.chCmd <- Command{
		Op:     opHGet,
		Key:    key,
		SubKey: subKey,
		Data:   data,
		Ret:    ret,
	}
	res := <-ret
	return res.Data.([]byte), res.Err
}

func (s *Store) HGetAll(key string) (m map[string][]byte, err error) {
	ret := make(chan Result, 1)
	s.chCmd <- Command{
		Op:  opHGetAll,
		Key: key,
		Ret: ret,
	}
	res := <-ret
	return res.Data.(map[string][]byte), res.Err
}

func (s *Store) HDelAll(key string) (err error) {
	ret := make(chan Result, 1)
	s.chCmd <- Command{
		Op:  opHDelAll,
		Key: key,
		Ret: ret,
	}
	res := <-ret
	return res.Err
}

func (s *Store) SyncExe(cmd Command) Result {
	cmd.Ret = make(chan Result, 1)
	s.chCmd <- cmd
	return <-cmd.Ret
}

func (s *Store) set(key string, data []byte) (err error) {
	_, ok := s.area[key]
	if ok {
		err = appErr.ErrStoreRecAlreadyExists
		return
	}
	s.area[key] = data
	return
}

func (s *Store) get(key string) (res []byte, err error) {
	iRes, ok := s.area[key]
	if !ok {
		err = appErr.ErrStoreRecNotFound
		return
	}
	res = iRes.([]byte)
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

func (s *Store) batch(data []byte) (err error) {
	var queued []Command
	err = json.Unmarshal(data, &queued)
	if err != nil {
		return
	}
	for _, cmd := range queued {
		res := s.handleCmd(cmd)
		if res.Err != nil {
			return res.Err
		}
	}
	return
}

// TODO: This function has serious performance problem. Potential optimizations include:
// 1) use dict tree
// 2) introduce scan, which would return only a part of data
func (s *Store) keys(keyPrefix string) (res [][]byte, err error) {
	for key, val := range s.area {
		if strings.HasPrefix(key, keyPrefix) {
			res = append(res, val.([]byte))
		}
	}
	return
}

// TODO: inconsistent with op set, which would report 'record already existed' error
func (s *Store) hset(key string, subKey string, data []byte) (err error) {
	iVal, ok := s.area[key]
	if !ok || iVal == nil {
		s.area[key] = make(map[string][]byte)
	}

	val := s.area[key].(map[string][]byte)
	val[subKey] = data
	return
}

func (s *Store) hget(key string, subKey string) (res []byte, err error) {
	iVal, ok := s.area[key]
	if !ok {
		err = appErr.ErrStoreRecNotFound
		return
	}
	val := iVal.(map[string][]byte)
	res, ok = val[subKey]
	if !ok {
		err = appErr.ErrStoreSubKeyNotFound
		return
	}
	return
}

func (s *Store) hgetall(key string) (res map[string][]byte, err error) {
	iVal, ok := s.area[key]
	if !ok {
		err = appErr.ErrStoreRecNotFound
		return
	}
	res = iVal.(map[string][]byte)
	return
}

func (s *Store) hdelall(key string) (err error) {
	_, ok := s.area[key]
	if !ok {
		err = appErr.ErrStoreRecNotFound
		return
	}
	delete(s.area, key)
	return
}
