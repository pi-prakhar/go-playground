package main

import (
	"fmt"
)

/*
FR :
1) Two Players take turn
2) dynamic board size 3x3 , 4x4, nxn
3) Players make move by specifying row and col : check for valid moves
4) Detect game status : win(row, col, diagonal), Draw( board is full and no more moves), Ongoing

NFR:
1) check for winners in O(1)

Bonus:
CLI
*/

/*
structs :
1) TicTacToe
   	variables - size int, board[char][char] n*n, rows[int] n, cols[int] n, diagonal int, anti-diagonal int, movesLeft n*n int, gameOver bool
   	public Methods:
    	StartGame()
    	MakeMove(player, row, col) status, error
     	DisplayBoard()
    private Methods:
    	CheckValidMove(row, col) bool
     	CheckWinner(row, col) bool
      	CheckGameOver(row, col) bool
       	updateBoard(player, row, col)
        displayMessage()
        getInput(player, *row, *col)
*/

type TicTacToe struct {
	size         int
	Board        [][]rune
	rows         []int
	cols         []int
	diagonal     int
	antiDiagonal int
	movesLeft    int
	gameOver     bool
}

func (t *TicTacToe) StartGame() {
	t.displayMessage()

	isPlayer1 := true

	for t.movesLeft > 0 {
		var row int
		var col int

		t.displayBoard()

		t.getInput(isPlayer1, &row, &col)

		// Check valid Move
		if !t.checkValidMove(row, col) {
			continue
		}

		// Make Move
		t.makeMove(isPlayer1, row, col)

		// Update Board
		t.updateBoard(isPlayer1, row, col)

		// Check Winner
		if t.checkWinner(row, col) {
			t.displayBoard()
			t.displayWinner(isPlayer1)
			return
		} else {
			t.movesLeft -= 1
			isPlayer1 = !isPlayer1
		}
	}

	t.displayBoard()
	t.displayDraw()
}

func (t *TicTacToe) updateBoard(isPlayer1 bool, row int, col int) {
	var marker rune
	if isPlayer1 {
		marker = 'x'
	} else {
		marker = 'o'
	}
	t.Board[row][col] = marker
}

func (t *TicTacToe) makeMove(isPlayer1 bool, row int, col int) {
	var count int
	if isPlayer1 {
		count = 1
	} else {
		count = -1
	}

	t.rows[row] += count
	t.cols[col] += count

	if row == col {
		t.diagonal += count
	}
	if row+col == t.size-1 {
		t.antiDiagonal += count
	}
}

func (t *TicTacToe) checkValidMove(row int, col int) bool {
	if row < t.size && row >= 0 && col < t.size && col >= 0 && t.Board[row][col] == '-' {
		return true
	}
	return false
}

func (t *TicTacToe) checkWinner(row int, col int) bool {
	if t.diagonal == t.size || t.diagonal == 0-t.size || t.antiDiagonal == t.size || t.antiDiagonal == 0-t.size || t.rows[row] == t.size || t.rows[row] == 0-t.size || t.cols[col] == t.size || t.cols[col] == 0-t.size {
		return true
	}
	return false
}

func (t *TicTacToe) getInput(isPlayer1 bool, row *int, col *int) {
	if isPlayer1 {
		fmt.Println("Player 1 : ")
	} else {
		fmt.Println("Player 2 : ")
	}
	fmt.Println("row : ")
	fmt.Scanf("%d", row)
	fmt.Println("col : ")
	fmt.Scanf("%d", col)
}

func (t *TicTacToe) displayMessage() {
	//Display Message
	fmt.Println("Player 1 : 'X', Player 2 : 'O'")
	fmt.Println("Player 1 will Start")
}

func (t *TicTacToe) displayBoard() {
	for _, row := range t.Board {
		for _, val := range row {
			fmt.Printf("%c ", val)
		}
		fmt.Println()
	}
}

func (t *TicTacToe) displayDraw() {
	fmt.Println("Match is Draw !!!")
}

func (t *TicTacToe) displayWinner(isPlayer1 bool) {
	if isPlayer1 {
		fmt.Println("Winner is Player 1")
	} else {
		fmt.Println("Winner is Player 2")
	}
}

// Factory for TicTacToe
func NewTicTacToe(size int) *TicTacToe {
	// init Board
	board := make([][]rune, size)
	for row := range size {
		board[row] = make([]rune, size)
		for col := range size {
			board[row][col] = '-'
		}
	}

	// init Tic Tac Toe
	ttt := TicTacToe{
		size:         size,
		Board:        board,
		rows:         make([]int, size),
		cols:         make([]int, size),
		diagonal:     0,
		antiDiagonal: 0,
		movesLeft:    size * size,
		gameOver:     false,
	}

	return &ttt
}

func main() {
	fmt.Println("Hello")
	ttt := NewTicTacToe(3)
	ttt.StartGame()
}
