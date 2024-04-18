package checkers

import (
	"fmt"
	"regexp"
)

func ValidatePassword(password string) bool {
	// Minimum 8 znaków, przynajmniej jedna litera i jedna cyfra
	// Go nie wspiera lookahead, więc używamy dwóch osobnych wyrażeń regularnych
	hasLetter := regexp.MustCompile(`[a-zA-Z0-9]`)
	hasProperLength := len(password) >= 8
	fmt.Printf("%v %v %v\n", hasLetter.MatchString(password), hasProperLength, password)
	return hasLetter.MatchString(password) && hasProperLength

}
