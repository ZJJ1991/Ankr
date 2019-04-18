package dataStructure



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