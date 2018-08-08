package helpers

import (
	"errors"
	"strconv"
)

var visaelectron map[string][]interface{}

type digits [6]int

func (d *digits) At(i int) int {
	return d[i-1]
}

//ReturningTypeCreditCard
func ReturningTypeCreditCard(num string) (Company, error) {
	ccLen := len(num)
	ccDigits := digits{}

	for i := 0; i < 6; i++ {
		if i < ccLen {
			ccDigits[i], _ = strconv.Atoi(num[:i+1])
		}
	}

	switch {
	case ccDigits.At(2) == 34 || ccDigits.At(2) == 37:
		return Company{"amex", "American Express"}, nil
	case ccDigits.At(4) == 5610 || (ccDigits.At(6) == 560221 && ccDigits.At(6) == 560225):
		return Company{"bankcard", "Bankcard"}, nil
	case ccDigits.At(2) == 62:
		return Company{"china-unionpay", "China UnionPay"}, nil
	case ccDigits.At(3) >= 300 && ccDigits.At(3) <= 305 && ccLen == 15:
		return Company{"diners-club-carte-blanche", "Diners Club Carte Blanche"}, nil
	case ccDigits.At(4) == 2014 || ccDigits.At(4) == 2149:
		return Company{"diners-club-enroute", "Diners Club enRoute"}, nil
	case ((ccDigits.At(3) >= 300 && ccDigits.At(3) <= 305) || ccDigits.At(3) == 309 || ccDigits.At(2) == 36 || ccDigits.At(2) == 38 || ccDigits.At(2) == 39) && ccLen <= 14:
		return Company{"diners-club-international", "Diners Club International"}, nil
	case ccDigits.At(4) == 6011 || (ccDigits.At(6) >= 622126 && ccDigits.At(6) <= 622925) || (ccDigits.At(3) >= 644 && ccDigits.At(3) <= 649) || ccDigits.At(2) == 65:
		return Company{"discover", "Discover"}, nil
	case ccDigits.At(3) == 636 && ccLen >= 16 && ccLen <= 19:
		return Company{"interpayment", "InterPayment"}, nil
	case ccDigits.At(3) >= 637 && ccDigits.At(3) <= 639 && ccLen == 16:
		return Company{"instapayment", "InstaPayment"}, nil
	case ccDigits.At(4) >= 3528 && ccDigits.At(4) <= 3589:
		return Company{"jcb", "JCB"}, nil
	case ccDigits.At(4) == 5018 || ccDigits.At(4) == 5020 || ccDigits.At(4) == 5038 || ccDigits.At(4) == 5612 || ccDigits.At(4) == 5893 || ccDigits.At(4) == 6304 || ccDigits.At(4) == 6759 || ccDigits.At(4) == 6761 || ccDigits.At(4) == 6762 || ccDigits.At(4) == 6763 || num[:3] == "0604" || ccDigits.At(4) == 6390:
		return Company{"maestro", "Maestro"}, nil
	case ccDigits.At(4) == 5019:
		return Company{"dankort", "Dankort"}, nil
	case ccDigits.At(2) >= 51 && ccDigits.At(2) <= 55:
		return Company{"mastercard", "MasterCard"}, nil
	case ccDigits.At(4) == 4026 || ccDigits.At(6) == 417500 || ccDigits.At(4) == 4405 || ccDigits.At(4) == 4508 || ccDigits.At(4) == 4844 || ccDigits.At(4) == 4913 || ccDigits.At(4) == 4917:
		return Company{"visa-electron", "Visa Electron"}, nil
	case ccDigits.At(1) == 4:
		return Company{"visa", "Visa"}, nil
	default:
		return Company{"", ""}, errors.New("Unknown credit card method.")
	}
}
