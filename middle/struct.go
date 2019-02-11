package middle

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var allowedApi = map[string]bool{"teams": true, "purse": true, "players": true}
