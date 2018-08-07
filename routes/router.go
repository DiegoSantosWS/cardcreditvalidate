package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DiegoSantosWS/cardcreditvalidate/helpers"
	"github.com/DiegoSantosWS/cardcreditvalidate/models"
	"github.com/gorilla/mux"
)

func Routers() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!!\r\n")
	})

	r.HandleFunc("/validate-cpf/{number}", func(w http.ResponseWriter, r *http.Request) {
		cod := mux.Vars(r)

		var retornoJS []helpers.Message
		m := helpers.Message{}
		if r.Method == "GET" {
			msg, err := models.ValidateCPF(cod["number"])
			if err != nil {
				m.Msg = msg
				m.Err = true
			} else {
				m.Msg = msg
			}
			retornoJS = append(retornoJS, m)

			retornoJSON, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(retornoJSON)
		}
	})

	r.HandleFunc("/validate-cnpj/{number}", func(w http.ResponseWriter, r *http.Request) {
		cod := mux.Vars(r)

		var retornoJS []helpers.Message
		m := helpers.Message{}
		if r.Method == "GET" {
			msg, err := models.ValidateCNPJ(cod["number"])
			if err != nil {
				m.Msg = msg
				m.Err = true
			} else {
				m.Msg = msg
			}
			retornoJS = append(retornoJS, m)

			retornoJSON, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(retornoJSON)
		}
	})

	r.HandleFunc("/validate-credit-card/{number}", func(w http.ResponseWriter, r *http.Request) {
		cod := mux.Vars(r)

		var retornoJS []helpers.Message
		m := helpers.Message{}
		if r.Method == "GET" {
			msg, errs := models.CardValidate(cod["number"])
			if errs == false {
				m.Msg = msg
				m.Err = true
			} else {
				m.Msg = msg
			}
			retornoJS = append(retornoJS, m)

			retornoJSON, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(retornoJSON)
		}
	})
	helpers.Runn(r)
}
