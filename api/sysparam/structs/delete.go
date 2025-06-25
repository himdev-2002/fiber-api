package structs

type DeleteSysParamsRequest struct {
	ID string `uri:"id" binding:"required" validate:"required,min=1,number"`
}
