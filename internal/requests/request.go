package requests

// user

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GetUser struct {
	Name string `json:"name" uri:"name" validate:"required"`
}

type DeleteUser struct {
	Name string `json:"name" uri:"name"  validate:"required"`
}

// role

type CreateRole struct {
	Name string `json:"name" validate:"required"`
}

type GetRole struct {
	Name string `json:"name" uri:"name" validate:"required"`
}

type DeleteRole struct {
	Name string `json:"name" uri:"name"  validate:"required"`
}

// user-role

type BindUserRole struct {
	UserName string `json:"userName" validate:"required"`
	RoleName string `json:"roleName" validate:"required"`
}

// authn

type Authenticate struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Invalidate struct {
	Token string `json:"token" validate:"required"`
}
