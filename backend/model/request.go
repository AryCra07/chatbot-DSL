package model

type RegisterLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ChatRequest struct {
}
