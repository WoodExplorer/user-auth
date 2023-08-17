package responses

const (
	CodeOK    = 0
	CodeError = 1
)

type Wrapper struct {
	Code int         `json:"code"`
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

// authn

type Authenticate struct {
	Token string `json:"token"`
}

// authz

type CheckRole struct {
	Ok bool `json:"ok"`
}

type UserRoles struct {
	Roles []Role `json:"roles"`
}
