package checkers

func CheckCredsLoginPass(login string, pass string) bool {
	return ValidatePassword(pass) && ValidateUsername(login)
}

func CheckCredsLoginEmail(login string, email string) bool {
	return ValidateEmail(email) && ValidateUsername(login)
}

func CheckCredsAll(pass string, email string, login string) bool {
	return ValidateEmail(email) && ValidatePassword(pass) && ValidateUsername(login)
}
