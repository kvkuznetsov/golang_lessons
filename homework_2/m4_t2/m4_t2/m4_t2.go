/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок
    
In: 1
Out: 1 бутылка
    
In: 22
Out: 22 бутылки
*/

package main

import (
    "fmt"
    "example.com/bottle"
)

func main() {
	var count int

    fmt.Println("Введите количество бутылок, от 0 до 200")
    fmt.Scanf("%d\n", &count)

    if count < 0 || count > 200 {
        fmt.Println("Некорретный размер, допустимый диапазон (0 до 200)")
        return
    }
	
    bottle.CountFmtPrint(count)
}