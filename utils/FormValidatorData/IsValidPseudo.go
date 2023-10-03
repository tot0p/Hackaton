package FormValidatorData

import "regexp"

func IsValidPseudo(pseudo string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]{3,12}$`).MatchString(pseudo)
}
