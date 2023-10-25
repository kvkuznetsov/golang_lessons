/* 
Задача №1
Написать функцию, которая расшифрует строку. 
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????' 
*/

package main

import (
    "fmt"
    "example.com/deencrypter"
)

func main() {
	var code string

	fmt.Println("Введите строку для шифрования")
    fmt.Scanf("%s\n", &code)

    fmt.Printf("Расшифрованная строка = %v\n", deencrypter.Deencrypt(code, 2))
}