package poker

import (
	"fmt"
	"io"
	"os"
	"time"
)

// BlindAlerter schedules alerts for blind amounts.
// `to io.Writer` takes a destination for the alerts so that we can re-use it in our webserver.
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// BlindAlerterFunc allows you to implement BlindAlerter with a function
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter.
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

// StdOutAlerter will schedule alerts and print them to os.Stdout.
func StdOutAlerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}

// Alerter will schedule alerts and print them to "to".
func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
