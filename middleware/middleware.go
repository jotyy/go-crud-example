package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jotyy/go-crud-example/auth"
	"github.com/jotyy/go-crud-example/response"
)

func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := auth.ExtractToken(c)
		if token == "" {
			code = 40001
		} else {
			err := auth.TokenValid(token)
			if err != nil {
				c.Abort()
				code = 40004
			}
		}

		if code != 200 {
			c.JSON(200, response.Response{
				Code: 40004,
				Msg:  "token认证失败",
				Data: data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
