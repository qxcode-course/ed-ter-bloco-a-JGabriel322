package main

import (
	"bufio"
	"fmt"
	"os"
)
	//_, _, _ = grid, startPos, endPos
	//return false

type Pos struct{
	l int
	c int
}

func isValid(grid [][]rune, l, c int) bool {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) {
		return false
	}

	return grid[l][c] == ' ' || grid[l][c] == 'F'
}

func search(grid [][]rune, atual Pos, end Pos) bool{
	if atual.l == end.l && atual.c == end.c {
		grid[atual.l][atual.c] = '.'
		return true
	}

	if !isValid(grid, atual.l, atual.c){
		return false
	}

	grid[atual.l][atual.c] = '.'

	if search(grid, Pos{atual.l+1, atual.c}, end) || 
		search(grid, Pos{atual.l, atual.c-1}, end) ||
		search(grid, Pos{atual.l, atual.c+1}, end) ||
		search(grid, Pos{atual.l-1, atual.c}, end) {
			return true
		}
	
	grid[atual.l][atual.c] = ' '
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
