package models

import (
	"errors"
	"log"
	"strconv"

	"github.com/DiegoSantosWS/cardcreditvalidate/helpers"
)

func ValidateCPF(cpf string) (string, error) {
	doc, err := helpers.NumbersCPF(cpf)
	if err != nil {
		log.Fatal("ERRO: THE NUMBER NOT AN IS STRING")
		return "", err
	}

	var eq bool
	var dig string
	for _, val := range doc {
		if len(dig) == 0 {
			dig = string(val)
		}

		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return "Erro 1 na validacao", err
	}
	i := 10
	sum := 0
	for index := 0; index < len(doc)-2; index++ {
		pos, _ := strconv.Atoi(string(doc[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(doc[9]))
	if mod != digit1 {
		return "Erro 2 na validacao", err
	}

	i = 11
	sum = 0
	for index := 0; index < len(doc)-1; index++ {
		pos, _ := strconv.Atoi(string(doc[index]))
		sum += pos * i
		i--
	}

	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(doc[10]))
	if mod != digit2 {
		return "Erro 3 na validacao", err
	}
	return "Cpf Valído", nil
}

func ValidateCNPJ(cnpj string) (string, error) {
	cnpj, err := helpers.NumbersCNPJ(cnpj)
	if err != nil {
		log.Fatal("ERRO: THE NUMBER NOT AN IS STRING")
		return "erro 1", err
	}

	algs := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var algProdCpfDig1 = make([]int, 12, 12)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))
		sumTmp := val * intParsed
		algProdCpfDig1[key] = sumTmp
	}
	sum := 0
	for _, val := range algProdCpfDig1 {
		sum += val
	}
	digit1 := sum % 11
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}
	char12, _ := strconv.Atoi(string(cnpj[12]))
	if char12 != digit1 {
		return "erro 2", errors.New("CNPJ inválido")
	}
	algs = append([]int{6}, algs...)

	var algProdCpfDig2 = make([]int, 13, 13)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))

		sumTmp := val * intParsed
		algProdCpfDig2[key] = sumTmp
	}
	sum = 0
	for _, val := range algProdCpfDig2 {
		sum += val
	}

	digit2 := sum % 11
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}
	char13, _ := strconv.Atoi(string(cnpj[13]))
	if char13 != digit2 {
		return "erro 3", errors.New("CNPJ inválido")
	}

	return "CNPJ Valido", nil
}
