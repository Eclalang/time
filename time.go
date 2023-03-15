package time

import (
	"strings"
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

// Strftime returns a string representation of a date according to a specified format
func Strftime(format, date string) string {
	result := ""
	for index := 0; index < len(format); index = 0 {
		i := strings.Index(format, "%")
		if i == -1 {
			break
		}
		result += format[:i]

		switch format[i+1] {
		case 'd':
			result += date[8:10]
		case 'm':
			result += date[5:7]
		case 'Y':
			result += date[:4]
		case 'y':
			result += date[2:4]
		case 'H':
			result += date[11:13]
		case 'M':
			result += date[14:16]
		case 'S':
			result += date[17:19]
		case '%':
			result += "%"
		}
		format = format[i+2:]
	}
	return result
}

// Timer waits for a specified number of seconds
func Timer(sec int) {
	t := time.NewTimer(time.Duration(sec) * time.Second)
	<-t.C
}
