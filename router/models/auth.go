package models

type AuthRequest struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type AuthResponse struct {
	Msg string `form:"msg" json:"msg" xml:"user"  binding:"required"`
}
