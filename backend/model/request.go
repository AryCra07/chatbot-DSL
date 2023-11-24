package model

type RegisterLoginRequest struct {
	Data struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	} `json:"data"`
}

type MessageRequest struct {
	Data struct {
		Input string `json:"input"`
	} `json:"data"`
}

type TimerRequest struct {
	Data struct {
		NowTime string `json:"now_time"`
	}
}
