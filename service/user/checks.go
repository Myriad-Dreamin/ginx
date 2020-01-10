package userservice

import "unicode"

func CheckStrongPassword(pwd string) string {
	return ""
}

func CheckPhone(phone string) string {
	return ""
}

func SumDigit(s string) int {
	var t int
	for i := range s {
		if unicode.IsDigit(rune(s[i])) {
			t++
		}
	}
	return t
}

func SumLower(s string) int {
	var t int
	for i := range s {
		if unicode.IsLower(rune(s[i])) {
			t++
		}
	}
	return t
}

func SumUpper(s string) int {
	var t int
	for i := range s {
		if unicode.IsUpper(rune(s[i])) {
			t++
		}
	}
	return t
}
