package router

import (
	"github.com/WoodExplorer/user-auth/internal/requests"
	"github.com/gin-gonic/gin"
)

func (r Router) applyToken(c *gin.Context) (_ interface{}, err error) {
	var p requests.Authenticate
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	res, err := r.authnSvc.Authenticate(c, p)
	if err != nil {
		return
	}

	return res, nil
}

func (r Router) invalidateToken(c *gin.Context) (_ interface{}, err error) {
	var p requests.Invalidate
	if err = bindAndValidate(c, &p); err != nil {
		return
	}

	err = r.authnSvc.Invalidate(c, p)
	if err != nil {
		return
	}

	return
}
