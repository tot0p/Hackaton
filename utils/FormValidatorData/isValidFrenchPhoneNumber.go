package FormValidatorData

import "regexp"

// IsValidFrenchPhoneNumber Function to check if a french phone number is valid with this regex: ^(0|\+33|0033)[1-9](\d{2}){4}$
func IsValidFrenchPhoneNumber(phoneNumber string) bool {
	return regexp.MustCompile(`^(0|\+33|0033)[1-9](\d{2}){4}$`).MatchString(phoneNumber)
}
