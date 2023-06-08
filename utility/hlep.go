package utility

//type ResponseError struct {
//	Error string `json:"error"`
//	Code  int
//}

var MySigningKey = []byte("my_bl^%%^og_84775M")

var List = []string{"0", "1"}

func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
