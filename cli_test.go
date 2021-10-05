package go_http_application_with_tdd_test

import (
	go_http_application_with_tdd "github.com/crazyhl/go-http-application-with-tdd"
	"strings"
	"testing"
	"time"
)

func TestCli(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &go_http_application_with_tdd.StubPlayerScore{}
		dummySpyAlerter := &SpyBlindAlerter{}
		cli := go_http_application_with_tdd.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		go_http_application_with_tdd.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &go_http_application_with_tdd.StubPlayerScore{}
		dummySpyAlerter := &SpyBlindAlerter{}
		cli := go_http_application_with_tdd.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		go_http_application_with_tdd.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins \n")
		playerStore := &go_http_application_with_tdd.StubPlayerScore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := go_http_application_with_tdd.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatalf("expected a blind alter to be scheduled")
		}
	})
}

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{scheduledAt: duration, amount: amount})
}
