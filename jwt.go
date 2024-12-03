package server

import (
	"github.com/gin-gonic/gin"
)

const (
	UserID      = "user_id"
	AccessToken = "access_token"
)

func AccessTokenHandle(secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader(AccessToken)
		if accessToken == "" {
			c.Error(AuthException())
			c.Abort()
			return
		}
		accessTokenClaims := UserClaims{}
		err := accessTokenClaims.ParseToken(accessToken, secret)
		if err != nil {
			c.Error(AuthException())
			c.Abort()
			return
		}
		c.Set(UserID, *accessTokenClaims.UID)
		c.Next()
	}
}
