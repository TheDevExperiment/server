package models

type AuthRequest struct {
	UserId      string `form:"userId" json:"userId" xml:"userId"  binding:"required"`
	SecretToken string `form:"secretToken" json:"secretToken" xml:"secretToken" binding:"required"`
	IsGuest     bool   `form:"isGuest" json:"isGuest" xml:"isGuest" binding:"required"`
}

type AuthResponse struct {
	Message   string      `form:"message" json:"message" xml:"message"  binding:"required"`
	ErrorCode string      `form:"errorCode" json:"errorCode" xml:"errorCode" binding:"required"`
	Data      interface{} `form:"data" json:"data"`
}
