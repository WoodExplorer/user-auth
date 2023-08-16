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
