package structs

type GetActiveUsersRequest struct {
	UsernameLike  string `json:"username_like" form:"username_like,omitempty" validate:"omitempty,alphanum,min=3,max=20"`
	EmailLike     string `json:"email_like" form:"email_like,omitempty" validate:"omitempty,alphanumunicode,min=3,max=100"`
	FirstNameLike string `json:"firstname_like" form:"firstname_like,omitempty" validate:"omitempty,alpha,min=3,max=50"`
	LastNameLike  string `json:"lastname_like" form:"lastname_like,omitempty" validate:"omitempty,alpha,min=3,max=50"`
}

type GetUserByIDRequest struct {
	ID string `uri:"id" binding:"required" validate:"required,min=3,max=30"`
}
