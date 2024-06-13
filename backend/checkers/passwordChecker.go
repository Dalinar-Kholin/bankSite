package checkers

import (
	"regexp"
)

func ValidatePassword(password string) bool {
	// Minimum 8 znaków, przynajmniej jedna litera i jedna cyfra
	// Go nie wspiera lookahead, więc używamy dwóch osobnych wyrażeń regularnych
	if len(password) < 4 || len(password) > 32 {
		return false
	}

	// Sprawdzenie obecności przynajmniej jednej cyfry
	hasDigit := regexp.MustCompile(`\d`)
	if !hasDigit.MatchString(password) {
		return false
	}

	// Sprawdzenie obecności przynajmniej jednej małej litery
	hasLower := regexp.MustCompile(`[a-z]`)
	if !hasLower.MatchString(password) {
		return false
	}

	// Sprawdzenie obecności przynajmniej jednej dużej litery
	hasUpper := regexp.MustCompile(`[A-Z]`)
	if !hasUpper.MatchString(password) {
		return false
	}

	// Jeśli wszystkie warunki są spełnione
	return true
}
