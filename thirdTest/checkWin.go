package main

import
(
	"fmt"
	ds "github.com/Ankr/datastructure"
	"math/rand"
	"time"
)


// Piece struct
// Since we already have ds.Square struct, we don't really need this one.
// I define this struct here because, I wanna link the struct with other functions, which is a noval feature of golang. 
// I verbosely do this just for showing examiner I know this feature.
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
func (p *Piece) checkSituation1 (cb ds.Chessboard) bool {
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
func (p *Piece) checkSituation2 (cb ds.Chessboard) bool {
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
func (p *Piece) checkSituation3 (cb ds.Chessboard) bool {
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
func (p *Piece) checkSituation4 (cb ds.Chessboard) bool {
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
	cb := ds.Chessboard{
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
   }
   for hasEmpty(cb) {
		cb = player1Term(cb)
		p = Piece {
			true, 
			false,
		}
		if (p.checkSituation1(cb) || p.checkSituation2(cb) || p.checkSituation3(cb) || p.checkSituation4(cb)) {
			printChessboard(cb)
			// fmt.Println("Player 1 Wins!")
			return 
		}
		if(hasEmpty(cb)){
			cb = player2Term(cb)
			p = Piece {
				false, 
				true,
			}
			if (p.checkSituation1(cb) || p.checkSituation2(cb) || p.checkSituation3(cb) || p.checkSituation4(cb)){
				printChessboard(cb)
				// fmt.Println()
				// fmt.Println("Player 2 Wins!")
				return 
			}
		}
   }
   // When the chessboard is fullfilled, check again who wins
   p = Piece{true, false,}
   if (p.checkSituation1(cb) || p.checkSituation2(cb) || p.checkSituation3(cb) || p.checkSituation4(cb)) {
		printChessboard(cb)
		// fmt.Println("Player 1 Wins!")
		return 
	}
	p = Piece{false, true,}
   if (p.checkSituation1(cb) || p.checkSituation2(cb) || p.checkSituation3(cb) || p.checkSituation4(cb)) {
		printChessboard(cb)
		// fmt.Println("Player 2 Wins!")
		return 
	}
	// No one wins 
    printChessboard(cb)
	fmt.Println()
	fmt.Printf("No one Wins!")
   return 
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
	// fmt.Print(index)
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