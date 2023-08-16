package router

import (
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/gin-gonic/gin"
)

func (r Router) bindUserRole(c *gin.Context) (_ interface{}, err error) {
	var p requests.BindUserRole
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	err = r.usrRoleSvc.Bind(c, p)
	if err != nil {
		return
	}

	return
}
