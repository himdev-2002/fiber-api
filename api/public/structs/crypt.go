package structs

type GenerateBcryptRequest struct {
	Password string `json:"password" form:"password" binding:"required"`
}
