package main

import
(
	"fmt"
	ds "github.com/Ankr/datastructure"
)



var square = ds.Square {
	ds.Empty,
	ds.Black,
	ds.White,
}

type direction struct{
	Pos7 [2]int
	Pos6 [2]int
	Pos5 [2]int
	Pos1 [2]int // the current piece
	Pos2 [2]int
	Pos3 [2]int
	Pos4 [2]int
}
func main(){
	playChess()
}

func check(cb *ds.Chessboard, d direction, piece int) bool{

	if (checkRange(d.Pos1, d.Pos2, d.Pos3, d.Pos4)) {
		if (cb[d.Pos2[0]][d.Pos2[1]] == piece &&
			cb[d.Pos3[0]][d.Pos3[1]] == piece &&
			cb[d.Pos4[0]][d.Pos4[1]] == piece){
			return true
		}
	}
	if (checkRange(d.Pos1, d.Pos5, d.Pos2, d.Pos3)) {
		if (cb[d.Pos2[0]][d.Pos2[1]] == piece &&
			cb[d.Pos3[0]][d.Pos3[1]] == piece && 
			cb[d.Pos5[0]][d.Pos5[1]] == piece){
			return true
		}
	}
	if (checkRange(d.Pos1, d.Pos6, d.Pos5, d.Pos2)) {
		if (cb[d.Pos2[0]][d.Pos2[1]] == piece && 
			cb[d.Pos6[0]][d.Pos6[1]] == piece && 
			cb[d.Pos5[0]][d.Pos5[1]] == piece){
			return true
		}
	}
	if (checkRange(d.Pos1, d.Pos7, d.Pos6, d.Pos5)) {
		if (cb[d.Pos7[0]][d.Pos7[1]] == piece && 
			cb[d.Pos6[0]][d.Pos6[1]] == piece && 
			cb[d.Pos5[0]][d.Pos5[1]] == piece){
			return true
		}
	}
	return false
}

func dir(pos [2]int, directionNumber int) func() direction {
	var d direction
	return func() direction {
		if (directionNumber == 1){ // direction1
			d.Pos7 = [2]int{pos[0]-3, pos[1]-3}
			d.Pos6 = [2]int{pos[0]-2, pos[1]-2}
			d.Pos5 = [2]int{pos[0]-1, pos[1]-1}
			d.Pos1 = pos
			d.Pos2 = [2]int{pos[0]+1, pos[1]+1}
			d.Pos3 = [2]int{pos[0]+2, pos[1]+2}
			d.Pos4 = [2]int{pos[0]+3, pos[1]+3}	
		}
		if (directionNumber == 2){ // direction2
			d.Pos7 = [2]int{pos[0]-3, pos[1]+3}
			d.Pos6 = [2]int{pos[0]-2, pos[1]+2}
			d.Pos5 = [2]int{pos[0]-1, pos[1]+1}
			d.Pos1 = pos
			d.Pos2 = [2]int{pos[0]+1, pos[1]-1}
			d.Pos3 = [2]int{pos[0]+2, pos[1]-2}
			d.Pos4 = [2]int{pos[0]+3, pos[1]-3}
		}
		if (directionNumber == 3){ // direction3
			d.Pos7 = [2]int{pos[0], pos[1]-3}
			d.Pos6 = [2]int{pos[0], pos[1]-2}
			d.Pos5 = [2]int{pos[0], pos[1]-1}
			d.Pos1 = pos
			d.Pos2 = [2]int{pos[0], pos[1]+1}
			d.Pos3 = [2]int{pos[0], pos[1]+2}
			d.Pos4 = [2]int{pos[0], pos[1]+3}
		}
		if (directionNumber == 4){ // direction4
			d.Pos7 = [2]int{pos[0]-3, pos[1]}
			d.Pos6 = [2]int{pos[0]-2, pos[1]}
			d.Pos5 = [2]int{pos[0]-1, pos[1]}
			d.Pos1 = pos
			d.Pos2 = [2]int{pos[0]+1, pos[1]}
			d.Pos3 = [2]int{pos[0]+2, pos[1]}
			d.Pos4 = [2]int{pos[0]+3, pos[1]}
		}
		return d
	}
}


func checkRange(pos1 [2]int, pos2 [2]int, pos3 [2]int, pos4 [2]int) bool{
	if (pos1[0]>=0 && pos1[0]<=5 && pos1[1]>=0 && pos1[1]<=6 && 
		pos2[0]>=0 && pos2[0]<=5 && pos2[1]>=0 && pos2[1]<=6 && 
		pos3[0]>=0 && pos3[0]<=5 && pos3[1]>=0 && pos3[1]<=6 && 
		pos4[0]>=0 && pos4[0]<=5 && pos4[1]>=0 && pos4[1]<=6) {
		return true
	}
	return false
}

// Functions below are implemented for the 2nd test.
func playChess() {
	cb := &ds.Chessboard{}
	for (cb.HasEmpty()){
		// player1's term
		fmt.Println("player1's term")
		position := cb.Play(square.Black)
		if (check(cb, dir(position, 1)(), square.Black)||
			check(cb, dir(position, 2)(), square.Black)||
			check(cb, dir(position, 3)(), square.Black)||
			check(cb, dir(position, 4)(), square.Black)){
			fmt.Printf("Player1 Wins, last position: [%d,%d]\n", position[0], position[1])
			cb.PrintChessboard()
			return
		}
		if (cb.HasEmpty()){
			// player2's term
			fmt.Println("player2's term")
			position = cb.Play(square.White)
			if (check(cb, dir(position, 1)(), square.White)||
				check(cb, dir(position, 2)(), square.White)||
				check(cb, dir(position, 3)(), square.White)||
				check(cb, dir(position, 4)(), square.White)){
				fmt.Printf("Player2 Wins, last position: [%d,%d]\n", position[0], position[1])
				cb.PrintChessboard()
				return
			}
		}
	}
}
