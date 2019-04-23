package main

import
(
	"fmt"
	ds "github.com/Ankr/datastructure"
)


// Piece struct
// Since we already have ds.Square struct, we don't really need this one.
// I define this struct here because, I wanna link the struct with other functions, which is a noval feature of golang. 
// I verbosely do this just for showing the examiner I know this feature.
type Piece struct{
	Black bool
	White bool
}

var square = ds.Square {
	ds.Empty,
	ds.Black,
	ds.White,
}


func main(){
	playChess()
}

// check if there exists one [i, j] that
// cb[i+1, j+1], cb[i+2, j+2], cb[i+3, j+3] combined with cb[i, j]
// are the same color
func (p *Piece) checkSituation1 (cb *ds.Chessboard) bool {
	color := 0;
	if p.White == true {
		color = square.White
	}
	if p.Black == true {
		color = square.Black
	}
	fmt.Printf("player:%v's term\n", color)

	for i := 0; i<3; i++ {
		for j := 0; j<4; j++ {
			if (cb[i][j] == cb[i+1][j+1] && cb[i][j] == cb[i+2][j+2] && cb[i][j] == cb[i+3][j+3] && cb[i][j] == color) {
				fmt.Printf("player %d wins with situation 1\n", color)
				return true
			}
		}
	}
	return false
}

// check if there exists one [i, j] that
// cb[i-1, j+1], cb[i-2, j+2], cb[i-3, j+3] combined with cb[i, j]
// are the same color
func (p *Piece) checkSituation2 (cb *ds.Chessboard) bool {
	color := 0;
	if p.White == true {
		color = square.White
	}
	if p.Black == true {
		color = square.Black
	}
	for i := 3; i<6; i++ {
		for j := 0; j<4; j++ {
			if (cb[i][j] == cb[i-1][j+1] && cb[i][j] == cb[i-2][j+2] && cb[i][j] == cb[i-3][j+3] && cb[i][j] == color) {
				fmt.Printf("player %d wins with situation 2\n", color)
				return true
			}
		}
	}
	return false
}

// check if there exists one [i, j] that
// cb[i, j+1], cb[i, j+2], cb[i, j+3] combined with cb[i, j]
// are the same color
func (p *Piece) checkSituation3 (cb *ds.Chessboard) bool {
	color := 0;
	if p.White == true {
		color = square.White
	}
	if p.Black == true {
		color = square.Black
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			if (cb[i][j] == cb[i][j+1] && cb[i][j] == cb[i][j+2] && cb[i][j] == cb[i][j+3] && cb[i][j] == color) {
				fmt.Printf("player %d wins with situation 3\n", color)
				return true
			}
		}
	}
	return false
}

// check if there exists one [i, j] that
// cb[i+1, j], cb[i+2, j], cb[i+3, j] combined with cb[i, j]
// are the same color
func (p *Piece) checkSituation4 (cb *ds.Chessboard) bool {
	color := 0;
	if p.White {
		color = square.White
	}
	if p.Black {
		color = square.Black
	}
	for j := 0; j < 7; j++ {
		for i := 0; i < 3; i++ {
			if (cb[i][j] == cb[i+1][j] && cb[i][j] == cb[i+2][j] && cb[i][j] == cb[i+3][j] && cb[i][j] == color) {
				fmt.Printf("player %d wins with situation 4\n", color)
				return true
			}
		}
	}
	return false
}

// Functions below are implemented for the 2nd test.
func playChess(){
	p := Piece	{
		false,
		false,
	}
	cb := &ds.Chessboard{}
   for cb.HasEmpty() {
		cb.Player1Term()
		p = Piece {
			true, 
			false,
		}
		if (p.checkSituation1(cb) || p.checkSituation2(cb) || p.checkSituation3(cb) || p.checkSituation4(cb)) {
			cb.PrintChessboard()
			// fmt.Println("Player 1 Wins!")
			return 
		}
		if(cb.HasEmpty()){
			cb.Player2Term()
			p = Piece {
				false, 
				true,
			}
			if (p.checkSituation1(cb) || p.checkSituation2(cb) || p.checkSituation3(cb) || p.checkSituation4(cb)){
				cb.PrintChessboard()
				// fmt.Println()
				// fmt.Println("Player 2 Wins!")
				return 
			}
		}
   }
   // When the chessboard is fullfilled, check again who wins
   p = Piece{true, false,}
   if (p.checkSituation1(cb) || p.checkSituation2(cb) || p.checkSituation3(cb) || p.checkSituation4(cb)) {
		cb.PrintChessboard()
		// fmt.Println("Player 1 Wins!")
		return
	}
	p = Piece{false, true,}
   if (p.checkSituation1(cb) || p.checkSituation2(cb) || p.checkSituation3(cb) || p.checkSituation4(cb)) {
		cb.PrintChessboard()
		// fmt.Println("Player 2 Wins!")
		return 
	}
	// No one wins
    cb.PrintChessboard()
	fmt.Println()
	fmt.Printf("No one Wins!")
   return 
}
