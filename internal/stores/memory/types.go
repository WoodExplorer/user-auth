package memory

const (
	opSet   = "set"
	opGet   = "get"
	opDel   = "del"
	opKeys  = "keys"
	opBatch = "batch"
)

type Result struct {
	Err  error
	Data interface{}
}

type Command struct {
	Op   string
	Key  string
	Data []byte
	Ret  chan Result `json:"-"`
}
