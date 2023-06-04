package utility

type ResponseError struct {
	Error string `json:"error"`
	Code  int
}

var MySigningKey string = "my_bl^%%^og_84775M"
