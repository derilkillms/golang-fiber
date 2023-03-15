package request

type UserCreateRequest struct {
	Name    string `json:"name" form:"name" validate:"required"`
	Address string `json:"address" form:"address"`
	Phone   string `json:"phone" form:"phone"`
	Email   string `json:"email" form:"email" validate:"required,email"`
}
