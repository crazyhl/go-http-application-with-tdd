package go_http_application_with_tdd_test

import (
	go_http_application_with_tdd "github.com/crazyhl/go-http-application-with-tdd"
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &go_http_application_with_tdd.StubPlayerScore{}
		cli := go_http_application_with_tdd.NewCLI(playerStore, in)
		cli.PlayPoker()

		go_http_application_with_tdd.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &go_http_application_with_tdd.StubPlayerScore{}
		cli := go_http_application_with_tdd.NewCLI(playerStore, in)
		cli.PlayPoker()

		go_http_application_with_tdd.AssertPlayerWin(t, playerStore, "Cleo")
	})

}
