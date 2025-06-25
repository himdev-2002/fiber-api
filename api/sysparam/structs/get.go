package structs

type GetUserSysParamsRequest struct {
	UID string `uri:"uid" binding:"required" validate:"required,min=1,number"`
}
