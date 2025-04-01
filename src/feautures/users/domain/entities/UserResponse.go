package entities

type UserResponse struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Rol      string `json:"rol"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
