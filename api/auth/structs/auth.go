package structs

type SignInRequest struct {
	Account  string `json:"account" form:"account" binding:"required" validate:"required,alphanumunicode,min=3,max=100"`
	Password string `json:"password" form:"password" binding:"required" validate:"required,min=8,max=100"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required" validate:"required"`
}
