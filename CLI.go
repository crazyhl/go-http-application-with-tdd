package go_http_application_with_tdd

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	PlayerStore PlayerStore
	In          io.Reader
}

func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.In)
	reader.Scan()
	c.PlayerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
