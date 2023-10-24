/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
*/

package main

import (
    "fmt"
    "example.com/chess_board_printer"
)

func main() {
	var size int

    fmt.Println("Введите размер шахматной доски, от 0 до 20")
    fmt.Scanf("%d\n", &size)

    if size < 0 || size > 20 {
        fmt.Println("Некорретный размер, допустимый диапазон (0 до 20)")
        return
    }
	
    chess_board_printer.PrintMatrix(size)
}