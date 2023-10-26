/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и 
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример: 
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8 
*/

package main

import (
    "fmt"
    "example.com/password"
	"bufio"
	"os"
	"log"
	"strings"
)

func main() {
	statusChan := make(chan password.PswStatus)
	reader := bufio.NewReader(os.Stdin)

	var inputPswd string
	var maxAttempt int = 5

	for i := 0; i < maxAttempt; i++ {
		fmt.Printf("Введите пароль, осталось попыток %v\n", maxAttempt - i)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputPswd = strings.TrimSpace(input)
		go password.IsCorrectLen(inputPswd, statusChan)
		go password.IsContainDigit(inputPswd, statusChan)
		go password.IsContainLowercase(inputPswd, statusChan)
		go password.IsContainUppercase(inputPswd, statusChan)
		go password.IsContainSpecial(inputPswd, statusChan)
		var isCorrectFull bool = true
		for i := 0; i < 5; i++ {
			status := <- statusChan
			fmt.Printf("Проверка: %v, Статус: %v %v\n", status.CheckName, status.Status, status.StatusText)
			if !status.Status { isCorrectFull = false }
		}
		if (isCorrectFull) { break }
	}
}