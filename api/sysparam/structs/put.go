package structs

type PutUserSysParamsRequest struct {
	ID       string  `uri:"id" binding:"required" validate:"required,min=1,number"`
	CatID    uint64  `json:"cat_id" form:"cat_id" validate:"min=0,number"`
	SubCatID uint64  `json:"sub_cat_id" form:"_sub_cat_id" validate:"min=0,number"`
	Stat     uint8   `json:"stat" form:"stat" validate:"oneof=0 1"` // 0: inactive, 1: active
	IsSystem uint8   `json:"is_system" validate:"oneof=0 1"`
	Str1     string  `json:"_str_1" validate:"max=255"`
	Str2     string  `json:"_str_2" validate:"max=255"`
	Str3     string  `json:"_str_3" validate:"max=255"`
	Num1     float64 `json:"_num_1" validate:"min=number"`
	Num2     float64 `json:"_num_2" validate:"min=number"`
	Num3     float64 `json:"_num_3" validate:"min=number"`
	Bool1    uint8   `json:"_bool_1" validate:"number,oneof=0 1"`
	Bool2    uint8   `json:"_bool_2" validate:"number,oneof=0 1"`
	Bool3    uint8   `json:"_bool_3" validate:"number,oneof=0 1"`
}
