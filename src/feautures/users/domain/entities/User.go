package domain

type User struct {
	IDUSER      int32   `json:"id_user"`
	USERNAME    string   `json:"username"`
	PASSWORD    string   `json:"password"`
	EMAIL       string   `json:"email"`
}