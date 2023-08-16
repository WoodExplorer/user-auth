package responses

type Wrapper struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// user

type User struct {
	Name string `json:"name"`
}

type GetUser struct {
	User
}

type ListUsers struct {
	Items []User `json:"items"`
}

// role

type Role struct {
	Name string `json:"name"`
}

type GetRole struct {
	Role
}

type ListRoles struct {
	Items []Role `json:"items"`
}
