package main

import "fmt"

var board[8][8]int 
var location int = 0 // used as global variable to track the location in which the knight is moving

func main(){
	board:= make([][]int, 8)
	for i:=range board{
		board[i] = make([]int, 8)
	}	
	for j:=0; j<8; j++{
		for k:=0; k<8; k++{
			board[j][k] = 0
		}
	}
	InitTour()
}

// moves that are allowed
func legalMove(row int, colm int) bool{
	if row>=0 && colm>=0 && row<8 && colm<8 {
		return true
	} 
	return false
}
// The possible way in which Knight can move in the board
func walk(row int, colm int, pos int) bool {
	if board[row][colm] != 0 {
		return false
	}
	location++
	board[row][colm] = location
	if pos == 63{
		return true
	}
//There are possible 8 moves that a Knight can take
	if legalMove(row+2, colm+1) && walk(row+2, colm+1, pos+1){
		return true
	}
	if legalMove(row+2, colm-1) && walk(row+2, colm-1, pos+1){
		return true
	}
	if legalMove(row-2, colm+1) && walk(row-2, colm+1, pos+1){
		return true
	}
	if legalMove(row-2, colm-1) && walk(row-2, colm-1, pos+1){
		return true
	}
	if legalMove(row+1, colm+2) && walk(row+1, colm+2, pos+1){
		return true
	}
	if legalMove(row+1, colm-2) && walk(row+1, colm-2, pos+1){
		return true
	}
	if legalMove(row-1, colm+2) && walk(row-1, colm+2, pos+1){
		return true
	}
	if legalMove(row-1, colm-2) && walk(row-1, colm-2, pos+1){
		return true
	}
// If above condition fails then
	board[row][colm] = 0
	location--
	return false
}

func InitTour(){
	if walk(0, 0, 8){
		for i:=0; i<8; i++{
			for j:=0; j<len(board); j++{
				fmt.Printf("%3d", board[i][j])
			}
				fmt.Println()
		}	
	} else {
		fmt.Println("The Program Fails to work")
	}
}




