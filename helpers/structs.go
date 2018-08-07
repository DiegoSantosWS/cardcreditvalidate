package helpers

//Messege retorna messagem
type Message struct {
	Msg string `json:"msg"`
	Err bool   `json:"err,omitempty"`
}
