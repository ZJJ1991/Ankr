package dataStructure

import
(
	"fmt"
	"time"
	"math/rand"
)

// The chessboard square status.
const (
	Empty = 0
	Black = 1
	White = 2
)

// A Square status of the chessboard, implying 
// if there exists a piece on a square and if so 
// whether a black or a white piece is placed on that square.
type Square struct {
	Empty int
	Black int
	White int
}

// The Chessboard has 6 rows and 7 columns.
type Chessboard [6][7]int

// PrintChessboard prints the whole chessboard.
func (cb *Chessboard) PrintChessboard(){
	for _, rowcontent := range cb	{
		for _, rowpiece  := range rowcontent {
			switch rowpiece{
			case Empty:
				fmt.Printf("%v\t", ".")	
			case Black:
				fmt.Printf("%v\t", "x")
			case White:
				fmt.Printf("%v\t", "o")
			}
		}
		fmt.Println()
	}
}

// HasEmpty checks if the chessboard has empty square or not.
func (cb *Chessboard) HasEmpty() bool{
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


func (cb *Chessboard) Play(piece int) *Chessboard {
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

func (cb *Chessboard) Player1Term() *Chessboard {
	cb.Play(Black)
	return cb
}

func (cb *Chessboard) Player2Term() *Chessboard {
	cb.Play(White)
	return cb
}
