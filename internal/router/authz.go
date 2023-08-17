package router

import (
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/gin-gonic/gin"
)

func (r Router) checkRole(c *gin.Context) (_ interface{}, err error) {
	var p requests.CheckRole
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	res, err := r.authzSvc.CheckRole(c, p)
	if err != nil {
		return
	}

	return res, nil
}

func (r Router) getUserRoles(c *gin.Context) (_ interface{}, err error) {
	var p requests.UserRoles
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	res, err := r.authzSvc.GetUserRoles(c, p)
	if err != nil {
		return
	}

	return res, nil
}
