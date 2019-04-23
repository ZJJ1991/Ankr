package main

import(
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
   cb := &ds.Chessboard{}
   for cb.HasEmpty() {
		cb.Play(square.Black)
		if(cb.HasEmpty()){
			cb.Play(square.White)			
		}
   }
   cb.PrintChessboard()
}

