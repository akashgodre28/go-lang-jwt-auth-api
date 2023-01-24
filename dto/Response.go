package dto

type Response struct {
	Message string      `json:"message"`
	Status  int16       `json:"status"`
	Data    interface{} `json:"data"`
}
