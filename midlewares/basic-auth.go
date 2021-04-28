package midlewares

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"dhikrama": "creator",
	})
}
