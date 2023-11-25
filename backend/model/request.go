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
		LastTime int32 `json:"last_time"`
		NowTime  int32 `json:"now_time"`
	} `json:"data"`
}
