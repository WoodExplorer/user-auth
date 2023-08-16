package memory

const (
	opSet   = "set"
	opGet   = "get"
	opDel   = "del"
	opBatch = "batch"
)

type Result struct {
	Err  error
	Data []byte
}

type Command struct {
	Op   string
	Key  string
	Data []byte
	Ret  chan Result `json:"-"`
}
