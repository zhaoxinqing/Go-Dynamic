package lib

import "regexp"

// PasswordIsValid ... Check whether the password is valid
func PasswordIsValid(password string) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{8,16}$", password); !ok {
		return false
	}
	return true
}

// VerifyMobileFormat ...Phone number correctness check
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
