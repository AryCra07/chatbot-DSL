package model

type RegisterLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserInfo struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type HelloRequest struct {
	Name string `json:"name"`
}

type MessageRequest struct {
	Name  string `json:"name"`
	Input string `json:"input"`
}
