package responses

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` //interface supaya data yang masuk sini bebas bisa string, array, object, dll
}
