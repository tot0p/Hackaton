package FormValidatorData

import "strings"

// IsPasswordSecure Function to check if a password is secure , it must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit and one special character
func IsPasswordSecure(password string) bool {
	return len(password) >= 8 && strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") && strings.ContainsAny(password, "0123456789") && strings.ContainsAny(password, `!@#$%^&*()-_=+{}[]|\:;\"'<>,.?/`)
}
