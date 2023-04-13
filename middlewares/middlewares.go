package middlewares

import (
	"net/http"
	"strings"

	"github.com/TheDevExperiment/server/internal/utility/jwt"
	"github.com/TheDevExperiment/server/router/models/auth"
	"github.com/gin-gonic/gin"
)

const ContextKeyUserId = "UserId"
const ContextKeyIsGuest = "IsGuest"

/*
if token is valid, attaches userId to the context.
Note that this middleware simply verifies the signature only.
It won't check if the Subject (in claim) is still an active user/not

Those aspect can be handled by handler functions themselves for now.
*/
func JWTAuthMiddleware(c *gin.Context) {
	var res auth.Response
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		res.Message = "Authorization header is missing"
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}
	authHeaderParts := strings.Split(authHeader, "Bearer ")
	if len(authHeaderParts) != 2 {
		res.Message = "Authorization header has incorrect format"
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}
	bearerToken := authHeaderParts[1]
	tokenClaims, err := jwt.VerifyToken(bearerToken)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}
	if tokenClaims == nil {
		res.Message = "Token is not valid"
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}
	c.Set(ContextKeyUserId, (*tokenClaims).Subject)
	c.Set(ContextKeyIsGuest, (*tokenClaims).IsGuest)
	c.Next()
}
