package bottle

import "fmt"

func CountFmtPrint(count int) {
	var suffix string = "ок"
	var modHundred int = count % 100
	var modTens int = count % 10

	if modHundred < 10 || modHundred > 15 {
		if modTens == 1 {
			suffix = "ка"
		} else if modTens > 1 && modTens < 5 {
			suffix = "ки"
		}
	}
	fmt.Printf("%v бутыл%v\n", count, suffix)
}