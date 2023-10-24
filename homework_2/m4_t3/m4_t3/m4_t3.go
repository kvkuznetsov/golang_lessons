/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/

package main

import (
    "fmt"
    "example.com/triangle"
)

func main() {
	var x1, y1, x2, y2, x3, y3 int

	fmt.Println("Введите последовательно координаты x1, y1, x2, y2, x3, y3 (через пробел)")
	fmt.Scanf("%d %d %d %d %d %d\n", &x1, &y1, &x2, &y2, &x3, &y3)

	fmt.Printf("X = (%v, %v)\n", x1, y1)
	fmt.Printf("Y = (%v, %v)\n", x2, y2)
	fmt.Printf("Z = (%v, %v)\n", x3, y3)

    fmt.Printf("Is Triangle Exist %v\n", triangle.IsExist(x1, y1, x2, y2, x3, y3))
	fmt.Printf("Triangle Area %v\n", triangle.GetArea(x1, y1, x2, y2, x3, y3))
	fmt.Printf("Is Triangle Rectangular %v\n", triangle.IsRectangular(x1, y1, x2, y2, x3, y3))
}