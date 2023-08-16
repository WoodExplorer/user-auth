package responses

type Wrapper struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type GetUser struct {
	Name string `json:"name"`
}

type GetRole struct {
	Name string `json:"name"`
}
