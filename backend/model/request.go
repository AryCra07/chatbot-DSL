package model

import "backend/log"

type RegisterLoginRequest struct {
	Data struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	} `json:"data"`
}

type UserInfo struct {
	Data struct {
		Name  string `json:"name"`
		State string `json:"state"`
	} `json:"data"`
}

type HelloRequest struct {
	Data struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	} `json:"data"`
}

type MessageRequest struct {
	Data struct {
		Name  string `json:"name"`
		Input string `json:"input"`
	} `json:"data"`
}

func (m RegisterLoginRequest) Logger() {
	log.Info("info", m.Data.Name+" "+m.Data.Password)
}
