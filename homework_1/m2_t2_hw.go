/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

package main

import "fmt"

func main() {
	var inputNumber int
	var isNumberPolindrom bool = false
	
	fmt.Println("Введите четырехзначное число (1000 - 9999)")
	fmt.Scanf("%d\n", &inputNumber)

	if inputNumber < 1000 || inputNumber > 9999 {
		fmt.Println("Некорретное значение, допустимый диапазон значений (100 - 999)")
		return
	}

	var thousands int = inputNumber / 1000
	var hundreds int = (inputNumber - thousands * 1000) / 100
	var tens int = (inputNumber - thousands * 1000 - hundreds * 100) / 10
	var units int = inputNumber - thousands * 1000 - hundreds * 100 - tens * 10

	if (thousands * 10 + hundreds == units * 10 + tens){
		isNumberPolindrom = true
	}

	if isNumberPolindrom {
		fmt.Printf("Number %d is Polindrom\n", inputNumber)
	} else {
		fmt.Printf("Number %d is not Polindrom\n", inputNumber)
	}
	
}