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
	pos7 := d.Pos7
	pos6 := d.Pos6
	pos5 := d.Pos5
	pos1 := d.Pos1
	pos2 := d.Pos2
	pos3 := d.Pos3
	pos4 := d.Pos4
	if (checkRange(pos1, pos2, pos3, pos4)) {
		if (cb[pos2[0]][pos2[1]] == piece &&
			cb[pos3[0]][pos3[1]] == piece &&
			cb[pos4[0]][pos4[1]] == piece){
			return true
		}
	}
	if (checkRange(pos1, pos5, pos2, pos3)) {
		if (cb[pos2[0]][pos2[1]] == piece &&
			cb[pos3[0]][pos3[1]] == piece && 
			cb[pos5[0]][pos5[1]] == piece){
			return true
		}
	}
	if (checkRange(pos1, pos6, pos5, pos2)) {
		if (cb[pos2[0]][pos2[1]] == piece && 
			cb[pos6[0]][pos6[1]] == piece && 
			cb[pos5[0]][pos5[1]] == piece){
			return true
		}
	}
	if (checkRange(pos1, pos7, pos6, pos5)) {
		if (cb[pos7[0]][pos7[1]] == piece && 
			cb[pos6[0]][pos6[1]] == piece && 
			cb[pos5[0]][pos5[1]] == piece){
			return true
		}
	}
	return false


}

func (d direction) direction1(pos [2]int) direction {
		d.Pos7 = [2]int{pos[0]-3, pos[1]-3}
		d.Pos6 = [2]int{pos[0]-2, pos[1]-2}
		d.Pos5 = [2]int{pos[0]-1, pos[1]-1}
		d.Pos1 = pos
		d.Pos2 = [2]int{pos[0]+1, pos[1]+1}
		d.Pos3 = [2]int{pos[0]+2, pos[1]+2}
		d.Pos4 = [2]int{pos[0]+3, pos[1]+3}
	return d
}

func (d direction) direction2(pos [2]int) direction {
		d.Pos7 = [2]int{pos[0]-3, pos[1]+3}
		d.Pos6 = [2]int{pos[0]-2, pos[1]+2}
		d.Pos5 = [2]int{pos[0]-1, pos[1]+1}
		d.Pos1 = pos
		d.Pos2 = [2]int{pos[0]+1, pos[1]-1}
		d.Pos3 = [2]int{pos[0]+2, pos[1]-2}
		d.Pos4 = [2]int{pos[0]+3, pos[1]-3}
	return d
}

func (d direction) direction3(pos [2]int) direction {
		d.Pos7 = [2]int{pos[0], pos[1]-3}
		d.Pos6 = [2]int{pos[0], pos[1]-2}
		d.Pos5 = [2]int{pos[0], pos[1]-1}
		d.Pos1 = pos
		d.Pos2 = [2]int{pos[0], pos[1]+1}
		d.Pos3 = [2]int{pos[0], pos[1]+2}
		d.Pos4 = [2]int{pos[0], pos[1]+3}
	return d
}

func (d direction) direction4(pos [2]int) direction {
		d.Pos7 = [2]int{pos[0]-3, pos[1]}
		d.Pos6 = [2]int{pos[0]-2, pos[1]}
		d.Pos5 = [2]int{pos[0]-1, pos[1]}
		d.Pos1 = pos
		d.Pos2 = [2]int{pos[0]+1, pos[1]}
		d.Pos3 = [2]int{pos[0]+2, pos[1]}
		d.Pos4 = [2]int{pos[0]+3, pos[1]}
	return d
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
		var d direction
		if (check(cb, d.direction1(position), square.Black)){
			fmt.Printf("Player1 Wins, last position: [%d,%d]\n", position[0], position[1])
			cb.PrintChessboard()
			return
		}
		if (check(cb, d.direction2(position), square.Black)){
			fmt.Printf("Player1 Wins, last position: [%d,%d]\n", position[0], position[1])
			cb.PrintChessboard()
			return
		}
		if (check(cb, d.direction3(position), square.Black)){
			fmt.Printf("Player1 Wins, last position: [%d,%d]\n", position[0], position[1])
			cb.PrintChessboard()
			return
		}
		if (check(cb, d.direction4(position), square.Black)){
			fmt.Printf("Player1 Wins, last position: [%d,%d]\n", position[0], position[1])
			cb.PrintChessboard()
			return
		}
		if (cb.HasEmpty()){
			// player2's term
			fmt.Println("player2's term")
			position = cb.Play(square.White)
			if (check(cb, d.direction1(position), square.White)){
				fmt.Printf("Player2 Wins, last position: [%d,%d]\n", position[0], position[1])
				cb.PrintChessboard()
				return
			}
			if (check(cb, d.direction2(position), square.White)){
				fmt.Printf("Player2 Wins, last position: [%d,%d]\n", position[0], position[1])				
				cb.PrintChessboard()
				return
			}
			if (check(cb, d.direction3(position), square.White)){
				fmt.Printf("Player2 Wins, last position: [%d,%d]\n", position[0], position[1])
				cb.PrintChessboard()
				return
			}
			if (check(cb, d.direction4(position), square.White)){
				fmt.Printf("Player2 Wins, last position: [%d,%d]\n", position[0], position[1])
				cb.PrintChessboard()
				return
			}
		}
	}
}
