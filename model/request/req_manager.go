package request

type LoginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type ChangePasswdRequest struct {
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	NewPasswd string `json:"newPasswd"`
}
