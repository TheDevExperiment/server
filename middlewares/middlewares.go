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
	var res auth.AuthResponse
	bearerToken := strings.Split(
		c.Request.Header.Get("Authorization"),
		"Bearer ")[1]
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
