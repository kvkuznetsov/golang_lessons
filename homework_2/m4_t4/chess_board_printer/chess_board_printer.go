package chess_board_printer

import "fmt"

func PrintMatrix(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%v ", (i  + j) % 2)
		}
		fmt.Print("\n")
	}
}