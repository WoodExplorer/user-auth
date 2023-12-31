package memory

import (
	"encoding/json"
	"github.com/WoodExplorer/user-auth/internal/stores"
)

type Tx struct {
	store  *Store
	queued []Command
}

func NewTx(store *Store) stores.Store {
	var tx Tx
	tx.store = store
	return &tx
}

func (t *Tx) Start() {
	// note: should not be called
	return
}

func (t *Tx) Stop() {
	// note: should not be called
	return
}

func (t *Tx) Set(key string, data []byte) (err error) {
	ret := make(chan Result, 1)
	t.queued = append(t.queued, Command{
		Op:   opSet,
		Key:  key,
		Data: data,
		Ret:  ret,
	})
	return
}

func (t *Tx) Get(_ string) (res []byte, err error) {
	// TODO: current tx implementation cannot interleave custom-code, so this function does nothing
	return
}

func (t *Tx) Del(key string) (err error) {
	ret := make(chan Result, 1)
	t.queued = append(t.queued, Command{
		Op:  opDel,
		Key: key,
		Ret: ret,
	})
	return
}

func (t *Tx) Keys(_ string) (data [][]byte, err error) {
	// TODO: current tx implementation cannot interleave custom-code, so this function does nothing
	return
}

func (t *Tx) HSet(key string, subKey string, data []byte) (err error) {
	ret := make(chan Result, 1)
	t.queued = append(t.queued, Command{
		Op:     opHSet,
		Key:    key,
		SubKey: subKey,
		Data:   data,
		Ret:    ret,
	})
	return
}

func (t *Tx) HGet(_ string, _ string) (data []byte, err error) {
	// TODO: current tx implementation cannot interleave custom-code, so this function does nothing
	return
}

func (t *Tx) HGetAll(_ string) (m map[string][]byte, err error) {
	// TODO: current tx implementation cannot interleave custom-code, so this function does nothing
	return
}

func (t *Tx) HDelAll(key string) (err error) {
	ret := make(chan Result, 1)
	t.queued = append(t.queued, Command{
		Op:  opHDelAll,
		Key: key,
		Ret: ret,
	})
	return
}

func (t *Tx) BeginTx() stores.Store {
	return t
}

func (t *Tx) CommitTx() error {
	bytes, _ := json.Marshal(t.queued)
	cmd := Command{
		Op:   opBatch,
		Data: bytes,
	}
	res := t.store.SyncExe(cmd)
	return res.Err
}
