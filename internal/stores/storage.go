package stores

type Store interface {
	Set(key string, data []byte) (err error)
	SetNx(key string, data []byte) (ok bool, err error)
	Get(key string) (res []byte, err error)
	GetE(key string) (res []byte, err error)
	Del(key string) (err error)
	DelE(key string) (err error)
}
