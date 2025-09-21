package phoneNumber

import "strconv"

func IsValid(phoneNumber string) bool {
	//TODO Add regular expression
	if len(phoneNumber) != 11 {
		return false
	}

	if phoneNumber[0:2] != "09" {
		return false
	}

	if _, err := strconv.Atoi(phoneNumber); err == nil {
		return false
	}

	return true
}
