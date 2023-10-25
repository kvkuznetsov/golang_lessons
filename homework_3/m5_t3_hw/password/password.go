package password

import (
	"unicode/utf8"
	"strings"
	"fmt"
)

func IsCorrectLen(checkableStr string) bool {
	var pswdLen int = utf8.RuneCountInString(checkableStr)
	if pswdLen > 7 && pswdLen < 16 {
		return true
	}
	fmt.Println("Некорретная длина пароля, допустимый диапазон от 8(мин) до 15(макс)")
	return false
}

func IsContainDigit(checkableStr string) bool {
	var result bool = isContainText(checkableStr, "0123456789")
	if !result {
		fmt.Println("Плохой пароль, не содержит чисел")
	}
	return result
}

func IsContainLowercase(checkableStr string) bool {
	var result bool = isContainText(checkableStr, "abcdefghiklmnopqrstvxyz")
	if !result {
		fmt.Println("Плохой пароль, не содержит букв в нижнем регистре")
	}
	return result
}

func IsContainUppercase(checkableStr string) bool {
	var result bool = isContainText(checkableStr, "ABCDEFGHIKLMNOPQRSTVXYZ")
	if !result {
		fmt.Println("Плохой пароль, не содержит букв в верхнем регистре")
	}
	return result
}

func IsContainSpecial(checkableStr string) bool {
	var result bool = isContainText(checkableStr, "_!@#$%^&")
	if !result {
		fmt.Println("Плохой пароль, не содержит специальных символов")
	}
	return result
}

func isContainText(checkableStr string, validateStr string) bool {
	for _, validateSimbol := range validateStr {
		if strings.Contains(checkableStr, string(validateSimbol)) {
			return true
		}
	}
	return false
}
