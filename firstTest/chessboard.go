package main
import (
	"fmt"
	ds "github.com/Ankr/datastructure"
)

var square = ds.Square {
	ds.Empty,
	ds.Black,
	ds.White,
}

func main(){
	fmt.Printf("Empty Chessboard \n")
	// Every chessboard elements will be automatically set 0 value.
	Chessboard := &ds.Chessboard{}
	Chessboard.PrintChessboard()
	fmt.Println()
	Chessboard[0][3] = square.Black
	Chessboard[2][3] = square.White
	Chessboard[3][2] = square.Black
	Chessboard[3][5] = square.White
	fmt.Printf("Case2 Chessboard \n")
	Chessboard.PrintChessboard()
}