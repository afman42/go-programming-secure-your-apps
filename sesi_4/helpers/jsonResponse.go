package helpers

type JSONResult200 struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type JSONResult422 struct {
	Message string     `json:"message"`
	Code    int        `json:"code"`
	Status  string     `json:"status"`
	Errors  []ApiError `json:"errors"`
}

type JSONResult400 struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type JSONResult401 struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
