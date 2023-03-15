package time

import (
	"time"
)

// Date returns a string representation of a date
func Date(year, month, day, hour, min, sec int) string {
	return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC).String()
}

// Now returns a string representation of the current time
func Now() string {
	return time.Now().String()
}

// Sleep pauses the current goroutine for a specified number of seconds
func Sleep(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

// TODO : DOES NOT WORK
func Strftime(format, date string) string {
	t, _ := time.Parse(time.RFC822, date)
	help := t.Format(format)
	return help
}

// Timer waits for a specified number of seconds
func Timer(sec int) {
	t := time.NewTimer(time.Duration(sec) * time.Second)
	<-t.C
}
