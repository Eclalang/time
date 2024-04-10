package time

import (
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"strconv"
	"strings"
	"time"
)

var Variables = map[string]eclaType.Type{
	"SECONDS": eclaType.Int(1000),
	"MINUTES": eclaType.Int(60000),
	"HOURS":   eclaType.Int(3600000),
}

// Date returns a string representation of a date
func Date(year, month, day, hour, min, sec int) string {
	return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC).String()
}

// Now returns a string representation of the current time
func Now() string {
	return time.Now().String()
}

// ConvertRoman converts a number to a roman number
func ConvertRoman(str string) string {
	nb, _ := strconv.Atoi(str)
	result := ""
	for nb > 0 {
		if nb >= 1000 {
			result += "M"
			nb -= 1000
		} else if nb >= 900 {
			result += "CM"
			nb -= 900
		} else if nb >= 500 {
			result += "D"
			nb -= 500
		} else if nb >= 400 {
			result += "CD"
			nb -= 400
		} else if nb >= 100 {
			result += "C"
			nb -= 100
		} else if nb >= 90 {
			result += "XC"
			nb -= 90
		} else if nb >= 50 {
			result += "L"
			nb -= 50
		} else if nb >= 40 {
			result += "XL"
			nb -= 40
		} else if nb >= 10 {
			result += "X"
			nb -= 10
		} else if nb >= 9 {
			result += "IX"
			nb -= 9
		} else if nb >= 5 {
			result += "V"
			nb -= 5
		} else if nb >= 4 {
			result += "IV"
			nb -= 4
		} else if nb >= 1 {
			result += "I"
			nb -= 1
		}
	}
	return result
}

// Sleep pauses the current goroutine for a specified number of milliseconds
func Sleep[T int | float64](ms T) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// Strftime returns a string representation of a date according to a specified format
func Strftime(format, date string) string {
	result := ""
	for index := 0; index < len(format); index = 0 {
		i := strings.Index(format, "%")
		if i == -1 {
			return format
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
