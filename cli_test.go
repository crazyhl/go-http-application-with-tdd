package go_http_application_with_tdd

import "testing"

func TestCli(t *testing.T) {
	playerStore := &StubPlayerScore{}
	cli := &CLI{playerStore}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatalf("expected a win call but didn't get any")
	}
}
