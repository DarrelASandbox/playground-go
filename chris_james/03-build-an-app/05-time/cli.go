package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	PlayerPrompt         = "Please enter the number of players: "
	BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
)

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

// Switch to `bufio.Scanner` instead of a reader as it's now automatically wrapped at construction time
type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{in: bufio.NewScanner(in), out: out, game: game}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)
	cli.game.Start(numberOfPlayers)
	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)
	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
