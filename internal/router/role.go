package router

import (
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/gin-gonic/gin"
)

func (r Router) createRole(c *gin.Context) (_ interface{}, err error) {
	var p requests.CreateRole
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	err = r.roleSvc.Create(c, p)
	if err != nil {
		return
	}

	return
}

func (r Router) getRole(c *gin.Context) (_ interface{}, err error) {
	var p requests.GetRole
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	res, err := r.roleSvc.Get(c, p)
	if err != nil {
		return
	}

	return res, nil
}

func (r Router) listRoles(c *gin.Context) (_ interface{}, err error) {

	res, err := r.roleSvc.List(c)
	if err != nil {
		return
	}

	return res, nil
}

func (r Router) deleteRole(c *gin.Context) (_ interface{}, err error) {
	var p requests.DeleteRole
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	err = r.roleSvc.Delete(c, p)
	if err != nil {
		return
	}

	return
}
