package FormValidatorData

import "regexp"

func IsValidFrenchPhoneNumber(phoneNumber string) bool {
	return regexp.MustCompile(`^(0|\+33|0033)[1-9](\d{2}){4}$`).MatchString(phoneNumber)
}
