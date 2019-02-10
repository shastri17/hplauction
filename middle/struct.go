package middle

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var allowedApi = map[string]bool{"team": true, "purse": true, "player": true}
