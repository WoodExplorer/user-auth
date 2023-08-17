package router

import (
	"fmt"
	"github.com/WoodExplorer/user-auth/internal/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

var (
	ValidateErrorMessage = map[string]string{}
	validate             = validator.New()
)

func Wrapper(handler func(*gin.Context) (interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {

		data, err := handler(c)

		resp := responses.Wrapper{Data: data}
		if err != nil {
			resp.Code = responses.CodeError
			resp.Msg = err.Error()
			log.Error().Msgf("handler error: %+v", err)
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Code = responses.CodeOK
		c.JSON(http.StatusOK, resp)
	}
}

func bindAndValidate(c *gin.Context, param interface{}) (err error) {
	if err = bind(c, param); err != nil {
		err = errors.Wrapf(err, "bind error")
		return
	}

	err = validate.Struct(param)
	if err != nil {
		err = errors.Wrapf(err, "validate error")
		return
	}

	return
}

func bind(c *gin.Context, params interface{}) error {
	_ = c.ShouldBindQuery(params)
	_ = c.ShouldBindUri(params)
	if err := c.ShouldBind(params); err != nil {
		if fieldErr, ok := err.(validator.ValidationErrors); ok {
			var tagErrorMsg []string
			for _, v := range fieldErr {
				if _, has := ValidateErrorMessage[v.Tag()]; has {
					tagErrorMsg = append(tagErrorMsg, fmt.Sprintf(ValidateErrorMessage[v.Tag()], v.Field(), v.Value()))
				} else {
					tagErrorMsg = append(tagErrorMsg, err.Error())
				}
			}

			return errors.New(strings.Join(tagErrorMsg, ","))
		}
	}

	return nil
}
