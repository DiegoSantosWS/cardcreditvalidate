package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Runn open instance of the server
func Runn(r *mux.Router) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	running := fmt.Sprintf(":%s", port)
	fmt.Printf("Rodando em: http://localhost%s\r\n", running)
	if err := http.ListenAndServe(running, r); err != nil {
		log.Fatal("[ERROR] Erro ao levantar o servidor.", err)
	}
}
