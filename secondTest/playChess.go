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
		cb = player1Term(cb)
		if(cb.HasEmpty()){
			cb = player2Term(cb)
		}
   }
   cb.PrintChessboard()
}

func player1Term(cb *ds.Chessboard) *ds.Chessboard {
	cb.Play(square.Black)
	return cb
}

func player2Term(cb *ds.Chessboard) *ds.Chessboard {
	cb.Play(square.White)
	return cb
}
