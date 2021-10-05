package go_http_application_with_tdd

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func (c *CLI) PlayPoker() {
	userInput := c.readline()
	c.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readline() string {
	c.in.Scan()
	return c.in.Text()
}
