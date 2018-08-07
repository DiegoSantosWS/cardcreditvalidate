package helpers

import (
	"errors"
	"strings"
)

//NumbersCPF contert string
func NumbersCPF(params string) (string, error) {
	params = strings.Replace(params, ".", "", -1)
	params = strings.Replace(params, "-", "", -1)

	if len(params) != 11 {
		return "", errors.New("CPF inválido")
	}
	return params, nil
}

//NumbersCPF contert string
func NumbersCNPJ(params string) (string, error) {
	params = strings.Replace(params, ".", "", -1)
	params = strings.Replace(params, "-", "", -1)
	params = strings.Replace(params, "/", "", -1)

	if len(params) != 14 {
		return "erro 1", errors.New("CNPJ inválido")
	}
	return params, nil
}
