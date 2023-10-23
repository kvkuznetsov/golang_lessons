/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/


package main

import "fmt"

func main() {
    var distance int;
    var consumption int;

    const petrolCost = 48.0;
    const distancePart = 100.0;

    fmt.Println("Введите расстояние (50 - 10000 км)")
    fmt.Scanf("%d\n", &distance)

    if distance < 50 || distance > 10000 {
        fmt.Println("Некорретное значение расстояния, допустимый диапазон значений (50 - 10000 км)")
        return
    }

    fmt.Println("Введите расход в литрах (5-25 литров) на 100 км")
    fmt.Scanf("%d\n", &consumption)

    if consumption < 5 || consumption > 25 {
        fmt.Println("Некорретное значение расхода, допустимый диапазон значений (5-25 литров)")
        return
    }

    var distanceParts float64 = float64(distance) / distancePart
    var petrolCostOnPart float64 = float64(consumption) * petrolCost

    var tripCost = distanceParts * petrolCostOnPart

    fmt.Printf("Стоимость поездки в рублях %v\n", tripCost)

}

