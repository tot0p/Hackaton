package FormValidatorData

import "strings"

func IsPasswordSecure(password string) bool {
	return len(password) >= 8 && strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") && strings.ContainsAny(password, "0123456789") && strings.ContainsAny(password, `!@#$%^&*()-_=+{}[]|\:;\"'<>,.?/`)
}
