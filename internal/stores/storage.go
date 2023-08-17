package stores

type Store interface {
	Start()
	Stop()

	Set(key string, data []byte) (err error)
	Get(key string) (res []byte, err error)
	Del(key string) (err error)
	Keys(keyPrefix string) (data [][]byte, err error)

	HSet(key string, subKey string, data []byte) (err error)
	HGet(key string, subKey string) (data []byte, err error)
	HGetAll(key string) (m map[string][]byte, err error)
	HDelAll(key string) (err error)

	BeginTx() Store
	CommitTx() (err error)
}
