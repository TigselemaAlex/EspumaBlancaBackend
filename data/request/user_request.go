package request

type UserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=16" json:"password"`
	Roles    []uint `validate:"required" json:"roles"`
	FullName string `validate:"required" json:"full_name"`
}
