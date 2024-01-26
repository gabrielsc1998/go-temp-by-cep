package cep_validators

import "regexp"

func IsValidCep(cep string) bool {
	cepRegex := `^\d{5}-\d{3}$|^\d{8}$`
	match, _ := regexp.MatchString(cepRegex, cep)
	return match
}
