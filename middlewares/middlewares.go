package middlewares

import (
	"net/http"

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
var JWTAuthMiddleware gin.HandlerFunc = func(c *gin.Context) {
	var req auth.AuthRequest
	var res auth.AuthResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}
	if req.SecretToken == "" {
		res.Message = "SecretToken must not be an empty string"
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}
	tokenClaims, err := jwt.VerifyToken(req.SecretToken)
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
