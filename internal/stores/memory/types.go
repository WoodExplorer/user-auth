package memory

const (
	opSet     = "set"
	opGet     = "get"
	opDel     = "del"
	opKeys    = "keys"
	opBatch   = "batch"
	opHSet    = "hset"
	opHGet    = "hget"
	opHGetAll = "hgetAll"
)

type Result struct {
	Err  error
	Data interface{}
}

type Command struct {
	Op     string
	Key    string
	SubKey string
	Data   []byte
	Ret    chan Result `json:"-"` // TODO: private?
}
