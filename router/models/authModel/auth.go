package authModel

type AuthRequest struct {
	SecretToken string `form:"secretToken" json:"secretToken" xml:"secretToken" binding:"required"`
}

type AuthResponse struct {
	Message   string      `form:"message" json:"message" xml:"message"  binding:"required"`
	ErrorCode string      `form:"errorCode" json:"errorCode" xml:"errorCode" binding:"required"`
	Data      interface{} `form:"data" json:"data" json:"data"`
}
