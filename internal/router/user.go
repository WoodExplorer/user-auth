package router

import (
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/gin-gonic/gin"
)

func (r Router) createUser(c *gin.Context) (_ interface{}, err error) {
	var p requests.CreateUser
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	err = r.usrSvc.Create(c, p)
	if err != nil {
		return
	}

	return
}

func (r Router) getUser(c *gin.Context) (_ interface{}, err error) {
	var p requests.GetUser
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	res, err := r.usrSvc.Get(c, p)
	if err != nil {
		return
	}

	return res, nil
}

func (r Router) listUsers(c *gin.Context) (_ interface{}, err error) {

	res, err := r.usrSvc.List(c)
	if err != nil {
		return
	}

	return res, nil
}

func (r Router) deleteUser(c *gin.Context) (_ interface{}, err error) {
	var p requests.DeleteUser
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	err = r.usrSvc.Delete(c, p)
	if err != nil {
		return
	}

	return
}
