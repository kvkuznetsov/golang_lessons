package password

import (
	"unicode/utf8"
	"strings"
)

type PswStatus struct {
	CheckName	string
	Status		bool
	StatusText	string
}


func IsCorrectLen(checkableStr string, channel chan PswStatus) {
	var pswdLen int = utf8.RuneCountInString(checkableStr)
	var checkName string = "Длина пароля"
	if pswdLen > 7 && pswdLen < 16 {
		channel <- PswStatus{CheckName: checkName, Status: true, StatusText: ""}
	} else {
		channel <- PswStatus{CheckName: checkName, Status: false, StatusText: "Некорретная длина пароля, допустимый диапазон от 8(мин) до 15(макс)"}
	}
}

func IsContainDigit(checkableStr string, channel chan PswStatus) {
	var result bool = isContainText(checkableStr, "0123456789")
	var checkName string = "Наличие цифр"
	if !result {
		channel <- PswStatus{CheckName: checkName, Status: false, StatusText: "Плохой пароль, не содержит чисел"}
	} else {
		channel <- PswStatus{CheckName: checkName, Status: true, StatusText: ""}
	}
}

func IsContainLowercase(checkableStr string, channel chan PswStatus) {
	var result bool = isContainText(checkableStr, "abcdefghiklmnopqrstvxyz")
	var checkName string = "Наличие букв в нижнем регистре"
	if !result {
		channel <- PswStatus{CheckName: checkName, Status: false, StatusText: "Плохой пароль, не содержит букв в нижнем регистре"}
	} else {
		channel <- PswStatus{CheckName: checkName, Status: true, StatusText: ""}
	}
}

func IsContainUppercase(checkableStr string, channel chan PswStatus) {
	var result bool = isContainText(checkableStr, "ABCDEFGHIKLMNOPQRSTVXYZ")
	var checkName string = "Наличие букв в верхнем регистре"
	if !result {
		channel <- PswStatus{CheckName: checkName, Status: false, StatusText: "Плохой пароль, не содержит букв в верхнем регистре"}
	} else {
		channel <- PswStatus{CheckName: checkName, Status: true, StatusText: ""}
	}
}

func IsContainSpecial(checkableStr string, channel chan PswStatus) {
	var result bool = isContainText(checkableStr, "_!@#$%^&")
	var checkName string = "Наличие специальных символов"
	if !result {
		channel <- PswStatus{CheckName: checkName, Status: false, StatusText: "Плохой пароль, не содержит специальных символов"}
	} else {
		channel <- PswStatus{CheckName: checkName, Status: true, StatusText: ""}
	}
}

func isContainText(checkableStr string, validateStr string) bool {
	for _, validateSimbol := range validateStr {
		if strings.Contains(checkableStr, string(validateSimbol)) {
			return true
		}
	}
	return false
}
