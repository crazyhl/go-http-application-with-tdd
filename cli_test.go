package go_http_application_with_tdd_test

import (
	"fmt"
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

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &go_http_application_with_tdd.StubPlayerScore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := go_http_application_with_tdd.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
}

func assertScheduledAlert(t *testing.T, got, want scheduledAlert) {
	if got.amount != want.amount {
		t.Errorf("got amount %d, want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time of %v, want %v", got.at, want.at)
	}
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at: duration, amount: amount})
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}
