/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/

package main

import "fmt"

func main() {
	var inputNumber1 int
	var inputNumber2 int
	var inputNumber3 int

	fmt.Println("Введите первое число")
	fmt.Scanf("%d\n", &inputNumber1)

	fmt.Println("Введите второе число")
	fmt.Scanf("%d\n", &inputNumber2)

	fmt.Println("Введите третие число")
	fmt.Scanf("%d\n", &inputNumber3)

	var maxNumber int = inputNumber1;
	var midNumber int = inputNumber2;
	var minNumber int = inputNumber3;

	if midNumber > maxNumber  {
		tmp := maxNumber
		maxNumber = midNumber
		midNumber = tmp
	}

	if minNumber > maxNumber  {
		tmp := maxNumber
		maxNumber = minNumber
		minNumber = tmp
	}

	if minNumber > midNumber  {
		tmp := midNumber
		midNumber = minNumber
		minNumber = tmp
	}

	fmt.Printf("%d, %d, %d\n", minNumber, midNumber, maxNumber)

}