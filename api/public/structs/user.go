package structs

type UserRegisterRequest struct {
	Username        string `json:"username" form:"username" binding:"required" validate:"required,alphanum,min=3,max=20"`
	Password        string `json:"password" form:"password" binding:"required" validate:"required,min=8,max=100"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm" binding:"required" validate:"required,eqfield=Password"`
	Email           string `json:"email" form:"email" binding:"required" validate:"required,email,min=5,max=100"`
	FirstName       string `json:"first_name" form:"first_name" binding:"required" validate:"required,alpha,min=3,max=50"`
	LastName        string `json:"last_name" form:"last_name" validate:"omitempty,alpha,min=3,max=50"`
}

type ChangePasswordRequest struct {
	Username  string `json:"username" form:"username" binding:"required" validate:"required,alphanum,min=3,max=20"`
	Password  string `json:"password" form:"password" binding:"required" validate:"required,min=8,max=100"`
	SecretKey string `json:"secret_key" form:"secret_key" binding:"required" validate:"required"`
}
