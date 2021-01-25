package request

type Login struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type Register struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type ChangePasswd struct {
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	NewPasswd string `json:"newPasswd"`
}
