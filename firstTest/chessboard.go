package main
import(
	"fmt"
	ds "github.com/Ankr/datastructure"
)

var square = ds.Square {
	ds.Empty,
	ds.Black,
	ds.White,
}

func main(){
	 fmt.Printf("Empty Chessboard \n\n")
	 emptyChessboard := ds.Chessboard{
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
	}
	printChessboard(emptyChessboard)
	fmt.Println()
	fmt.Printf("Case2 Chessboard \n\n")
	Case2Chessboard := ds.Chessboard{
		{square.Empty, square.Empty, square.Empty, square.Black, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.White, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Black, square.Empty, square.Empty, square.White, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},
		{square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty, square.Empty},	}
	printChessboard(Case2Chessboard)

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


