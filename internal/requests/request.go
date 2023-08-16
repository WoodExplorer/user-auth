package requests

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
