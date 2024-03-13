package time

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	year, month, day, hour, min, sec := 2003, 6, 9, 14, 21, 42
	actual := Date(year, month, day, hour, min, sec)
	expected := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC).String()
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestNow(t *testing.T) {
	actual := Now()
	expected := time.Now().String()
	actual = actual[0:19]
	expected = expected[0:19]
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestConvertRoman(t *testing.T) {
	actual := ConvertRoman("1999")
	expected := "MCMXCIX"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual = ConvertRoman("658")
	expected = "DCLVIII"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual = ConvertRoman("444")
	expected = "CDXLIV"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual = ConvertRoman("10")
	expected = "X"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSleep(t *testing.T) {
	Sleep(2)
	Sleep(0.5)
}

func TestStrftime(t *testing.T) {
	date := Date(2003, 6, 9, 14, 21, 42)
	format := "Current time : %d/%m/%Y %H:%M:%S"
	actual := Strftime(format, date)
	expected := "Current time : 09/06/2003 14:21:42"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	format = "Current time : "
	actual = Strftime(format, date)
	expected = "Current time : "
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	format = "Current time : %y %%"
	actual = Strftime(format, date)
	expected = "Current time : 03 %"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
