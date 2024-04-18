package checkers

import "regexp"

func ValidateUsername(username string) bool {
	// Tylko litery, cyfry i podkreślenia, od 3 do 16 znaków
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)
	return re.MatchString(username)
}
