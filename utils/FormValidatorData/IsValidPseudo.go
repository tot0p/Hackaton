package FormValidatorData

import "regexp"

// IsValidPseudo Function to check if a pseudo is valid with this regex: ^[a-zA-Z0-9]{3,12}$ (3 to 12 characters)
func IsValidPseudo(pseudo string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]{3,12}$`).MatchString(pseudo)
}
