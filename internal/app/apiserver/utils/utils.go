package utils

import "strings"

func ValidationID(uuid string) bool {
	rs := []rune(uuid)

	if len(rs) != 36 {
		return false
	}

	return true
}

func ValidateTransaction(uuid string, operation string, amount float64) bool {
	if ! ValidationID(uuid) {
		return false
	}

	op := strings.ToUpper(operation)
	if op != "DEPOSIT" && op != "WITHDRAW" {
		return false
	}

	if amount < 0 {
		return false
	}
	return true
}

