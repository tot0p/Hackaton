package FormValidatorData

import "regexp"

func IsPasswordSecure(password string) bool {
	return regexp.MustCompile(`^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@#$%^&+=!]).{8,}$`).MatchString(password)
}
