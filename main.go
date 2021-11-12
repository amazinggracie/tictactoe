package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func printBoard(board map[int]string) {
	fmt.Printf("%s | %s | %s\n", board[7], board[8], board[9])
	fmt.Println("- + - + -")
	fmt.Printf("%s | %s | %s\n", board[4], board[5], board[6])
	fmt.Println("- + - + -")
	fmt.Printf("%s | %s | %s\n", board[1], board[2], board[3])
}

func checkForWin(board map[int]string, player string) string {
	// For every turn after turn 4, we can assume that if
	// a win condition has been met, the player that went last is the winner.
	switch {
	// Top row
	case board[7]+board[8]+board[9] == strings.Repeat(player, 3):
		return player

	// Middle row
	case board[4]+board[5]+board[6] == strings.Repeat(player, 3):
		return player

	// Bottom row
	case board[1]+board[2]+board[3] == strings.Repeat(player, 3):
		return player

	// Left column
	case board[7]+board[4]+board[1] == strings.Repeat(player, 3):
		return player

	// Middle column
	case board[8]+board[5]+board[2] == strings.Repeat(player, 3):
		return player

	// Right column
	case board[9]+board[6]+board[3] == strings.Repeat(player, 3):
		return player

	// Diagonal left
	case board[7]+board[5]+board[3] == strings.Repeat(player, 3):
		return player

	// Diagonal right
	case board[1]+board[2]+board[3] == strings.Repeat(player, 3):
		return player
	}
	return ""
}

func main() {
	playerTurn := "X"
	turnCount := 0

	board := map[int]string{
		7: " ", 8: " ", 9: " ",
		4: " ", 5: " ", 6: " ",
		1: " ", 2: " ", 3: " ",
	}

	for turnCount < 9 {
		printBoard(board)

		var move int
		fmt.Printf("It's your turn %s, which place will you choose?\n", playerTurn)
		_, err := fmt.Scan(&move)

		if err == io.EOF {
			os.Exit(0) // Silently exit if EOF is found
		}

		// If the input is not an integer
		if err != nil {
			fmt.Println("Please enter a valid number, 1-9.")
			continue
		}

		// Check if the move is valid
		if move < 1 || move > 9 {
			fmt.Println("Please enter a valid number, 1-9.")
			continue
		}

		// Check if the place is already taken
		if board[move] == " " {
			board[move] = playerTurn
			turnCount += 1
		} else {
			fmt.Println("That place was already chosen.")
			continue
		}

		// Check if player X or O has won, every move after 5 turns.
		// A player (X or O) cannot win until at least 5 turns have been taken.
		if turnCount >= 5 {
			if checkForWin(board, playerTurn) == "X" {
				fmt.Println("X won.")
				printBoard(board)
				break
			}
			if checkForWin(board, playerTurn) == "O" {
				fmt.Println("O won.")
				printBoard(board)
				break
			}
		}

		if turnCount == 9 && checkForWin(board, playerTurn) == "" {
			fmt.Println("It's a tie.")
			printBoard(board)
		}

		// Change the player after every move
		if playerTurn == "X" {
			playerTurn = "O"
		} else if playerTurn == "O" {
			playerTurn = "X"
		}
	}
}
