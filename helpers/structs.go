package helpers

//Messege retorna messagem
type Message struct {
	Msg  string  `json:"msg"`
	Err  bool    `json:"err,omitempty"`
	Comp Company `json:"comp,omitempty"`
}

type Company struct {
	Short, Long string
}
