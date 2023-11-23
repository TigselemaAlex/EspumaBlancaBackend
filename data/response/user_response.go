package response

type UserResponse struct {
	ID       uint           `json:"id"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	FullName string         `json:"full_name"`
	Roles    []RoleResponse `json:"roles"`
}
