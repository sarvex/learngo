package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	games := load()
	byID := indexByID(games)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > save   : exports the data to json and quits
  > quit   : quits

`)

		if !in.Scan() || !runCmd(in.Text(), games, byID) {
			break
		}
	}
}
