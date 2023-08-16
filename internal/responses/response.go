package responses

type Wrapper struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type User struct {
	Name string `json:"name"`
}

type GetUser struct {
	User
}

type ListUsers struct {
	Items []User `json:"items"`
}

type GetRole struct {
	Name string `json:"name"`
}
