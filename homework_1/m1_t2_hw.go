/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/

package main

import "fmt"

func main() {
	var inputNumber int

	fmt.Println("Введите трехзначное число (100 - 999)")
	fmt.Scanf("%d\n", &inputNumber)

	if inputNumber < 100 || inputNumber > 999 {
		fmt.Println("Некорретное значение, допустимый диапазон значений (100 - 999)")
		return
	}

	var hundreds int = inputNumber / 100
	var tens int = (inputNumber - hundreds*100) / 10
	var units int = inputNumber - hundreds*100 - tens*10

	fmt.Printf("Result = %v%v%v\n", units, tens, hundreds)

}
