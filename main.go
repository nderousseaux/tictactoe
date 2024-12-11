package main

import "fmt"

func win(board [3][3]uint8) uint8 {
	// If j1 wins -> return 1
	// If j2 wins -> return 2
	// If no one wins -> return 0

	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}

	// Check diagonals
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}

	return 0
}

func printBoard(board [3][3]uint8) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				fmt.Print(" ")
			}
			if board[i][j] == 1 {
				fmt.Print("X")
			}
			if board[i][j] == 2 {
				fmt.Print("O")
			}
			if j < 2 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("---------")
		}
	}
}

func main() {
	board := [3][3]uint8{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	turn := 1
	for win(board) == 0 {
		printBoard(board)
		fmt.Print("Player ", turn)
		fmt.Println(" : Enter row and column:")
		var row, col int
		fmt.Scan(&row, &col)
		board[row][col] = uint8(turn)
		if win(board) != 0 {
			break
		}
		if turn == 1 {
			turn = 2
		} else {
			turn = 1
		}
	}

	printBoard(board)

	if win(board) == 1 {
		fmt.Println("Player 1 wins!")
	} else {
		fmt.Println("Player 2 wins!")
	}
}