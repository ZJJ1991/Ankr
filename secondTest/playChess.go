package main

import(
	"fmt"
	"math/rand"
	"time"
	ds "github.com/Ankr/datastructure"
)


var square = ds.Square {
	ds.Empty,
	ds.Black,
	ds.White,
}


func main(){
	playChess()
}

// Functions below are implemented for the 2nd test.
func playChess(){
	cb := ds.Chessboard{
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
   }
   hasempty := true
   for hasempty {
		cb = player1Term(cb)
		if(hasEmpty(cb)){
			cb = player2Term(cb)
		}
		hasempty = hasEmpty(cb)
   }
   printChessboard(cb)
}

func player1Term(cb ds.Chessboard) ds.Chessboard {
	refereshedChessboard := play(cb, square.Black)
	return refereshedChessboard
}

func player2Term(cb ds.Chessboard) ds.Chessboard {
	refereshedChessboard := play(cb, square.White)
	return refereshedChessboard
}

func play(cb ds.Chessboard, piece int) ds.Chessboard{
	// randomly choose one column
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(7)
	// form bottom to top
	for i := 5; i >= 0; i-- {
		if ( cb[i][index] == 0 ){
			cb[i][index] = piece
			return cb
		}
	}
	return cb
}

func hasEmpty(cb ds.Chessboard) bool{
	// from top to bottom
	for _, rowContent := range cb {
		for _, component := range rowContent {
			if ( component == 0 ){
				return true
			}
		}
	}
	return false
}

func printChessboard(cb ds.Chessboard){
	for _, rowcontent := range cb	{
		for _, rowcomponent  := range rowcontent {
			switch rowcomponent{
			case 0:
				fmt.Printf("%v\t", ".")	
			case 1:
				fmt.Printf("%v\t", "x")
			case 2:
				fmt.Printf("%v\t", "o")
			}
		}
		fmt.Println()
	}
}