package model

type Manager struct {
	AM_MODEL
	Nickname string `json:"nickname"`
	Password string `json:"-"`
}
