package stores

type Store interface {
	Start()
	Stop()

	Set(key string, data []byte) (err error)
	Get(key string) (res []byte, err error)
	Del(key string) (err error)
	Keys(keyPrefix string) (data [][]byte, err error)

	BeginTx() Store
	CommitTx() (err error)
}
