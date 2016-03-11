package simpletime

import (
	"testing"
	"time"
)

func TestBeginningOfYear(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), 1, 1, 0, 0, 0, 0, input.Location())

	result := NewSimpleTime(input).BeginningOfYear()
	if result != expected {
		t.Errorf("TestBeginningOfYear Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestBeginningOfMonth(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), 1, 0, 0, 0, 0, input.Location())

	result := NewSimpleTime(input).BeginningOfMonth()
	if result != expected {
		t.Errorf("TestBeginningOfMonth Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestBeginningOfDay(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), input.Day(), 0, 0, 0, 0, input.Location())

	result := NewSimpleTime(input).BeginningOfDay()
	if result != expected {
		t.Errorf("TestBeginningOfDay Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestBeginningOfHour(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), input.Day(), input.Hour(), 0, 0, 0, input.Location())

	result := NewSimpleTime(input).BeginningOfHour()
	if result != expected {
		t.Errorf("TestBeginningOfHour Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestBeginningOfMinute(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), input.Day(), input.Hour(), input.Minute(), 0, 0, input.Location())

	result := NewSimpleTime(input).BeginningOfMinute()
	if result != expected {
		t.Errorf("TestBeginningOfMinute Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestBeginningOfSecond(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), input.Day(), input.Hour(), input.Minute(), input.Second(), 0, input.Location())

	result := NewSimpleTime(input).BeginningOfSecond()
	if result != expected {
		t.Errorf("TestBeginningOfSecond Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestEndOfYear(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), 12, 31, 23, 59, 59, 999999999, input.Location())

	result := NewSimpleTime(input).EndOfYear()
	if result != expected {
		t.Errorf("TestEndOfYear Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestEndOfMonth(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), 28, 23, 59, 59, 999999999, input.Location())

	result := NewSimpleTime(input).EndOfMonth()
	if result != expected {
		t.Errorf("TestEndOfMonth Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestEndOfDay(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), input.Day(), 23, 59, 59, 999999999, input.Location())

	result := NewSimpleTime(input).EndOfDay()
	if result != expected {
		t.Errorf("TestEndOfDay Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestEndOfHour(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), input.Day(), input.Hour(), 59, 59, 999999999, input.Location())

	result := NewSimpleTime(input).EndOfHour()
	if result != expected {
		t.Errorf("TestEndOfHour Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestEndOfMinute(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), input.Day(), input.Hour(), input.Minute(), 59, 999999999, input.Location())

	result := NewSimpleTime(input).EndOfMinute()
	if result != expected {
		t.Errorf("TestEndOfMinute Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestEndOfSecond(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(input.Year(), input.Month(), input.Day(), input.Hour(), input.Minute(), input.Second(), 999999999, input.Location())

	result := NewSimpleTime(input).EndOfSecond()
	if result != expected {
		t.Errorf("TestEndOfMinute Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestNextOccurrenceOfWeekday(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := input

	result := NewSimpleTime(input).NextOccurrenceOfWeekday(input.Weekday())
	if result != expected {
		t.Errorf("TestNextOccurrenceOfWeekday Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}

	expected = time.Date(2001, 2, 9, 0, 0, 0, 0, time.UTC)

	result = NewSimpleTime(input).NextOccurrenceOfWeekday(time.Weekday(5))
	if result != expected {
		t.Errorf("TestNextOccurrenceOfWeekday Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestNextSunday(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(2001, 2, 4, 0, 0, 0, 0, time.UTC)

	result := NewSimpleTime(input).NextSunday()
	if result != expected {
		t.Errorf("TestNextSunday Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestNextMonday(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(2001, 2, 5, 0, 0, 0, 0, time.UTC)

	result := NewSimpleTime(input).NextMonday()
	if result != expected {
		t.Errorf("TestNextMonday Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestNextTuesday(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(2001, 2, 6, 0, 0, 0, 0, time.UTC)

	result := NewSimpleTime(input).NextTuesday()
	if result != expected {
		t.Errorf("TestNextTuesday Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestNextWednesday(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(2001, 2, 7, 0, 0, 0, 0, time.UTC)

	result := NewSimpleTime(input).NextWednesday()
	if result != expected {
		t.Errorf("TestNextWednesday Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestNextThursday(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(2001, 2, 8, 0, 0, 0, 0, time.UTC)

	result := NewSimpleTime(input).NextThursday()
	if result != expected {
		t.Errorf("TestNTestNextThursdayextSunday Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestNextFriday(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(2001, 2, 9, 0, 0, 0, 0, time.UTC)

	result := NewSimpleTime(input).NextFriday()
	if result != expected {
		t.Errorf("TestNextFriday Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestNextSaturday(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := input

	result := NewSimpleTime(input).NextSaturday()
	if result != expected {
		t.Errorf("TestNextSaturday Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestAddDays(t *testing.T) {
	tests := []struct {
		input    time.Time
		expected time.Time
		days     int
	}{
		{
			time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC),
			time.Date(2001, 2, 103, 4, 5, 6, 7, time.UTC),
			100,
		},
		{
			time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC),
			time.Date(2001, 2, 1, 4, 5, 6, 7, time.UTC),
			-2,
		},
	}
	for i, test := range tests {
		result := NewSimpleTime(test.input).AddDays(test.days)
		if result != test.expected {
			t.Errorf("TestAddDays %d Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", i, test.input, test.expected, result)
		}

	}

}

func TestNextDay(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(2001, 2, 4, 4, 5, 6, 7, time.UTC)

	result := NewSimpleTime(input).NextDay()
	if result != expected {
		t.Errorf("TestNextDay Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

func TestPrevDay(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	expected := time.Date(2001, 2, 2, 4, 5, 6, 7, time.UTC)

	result := NewSimpleTime(input).PrevDay()
	if result != expected {
		t.Errorf("TestPrevDay Failed:\nInput=(%v),\nExpected=(%v),\nResult=(%v)", input, expected, result)
	}
}

// Test from subTests in time_test.go
// https://github.com/golang/go/blob/master/src/time/time_test.go
func TestSince(t *testing.T) {
	tests := []struct {
		t time.Time
		u time.Time
		d *SimpleDuration
	}{
		{time.Time{}, time.Time{}, NewSimpleDuration(0)},
		{time.Date(2009, 11, 23, 0, 0, 0, 1, time.UTC), time.Date(2009, 11, 23, 0, 0, 0, 0, time.UTC), NewSimpleDuration(1)},
		{time.Date(2009, 11, 23, 0, 0, 0, 0, time.UTC), time.Date(2009, 11, 24, 0, 0, 0, 0, time.UTC), NewSimpleDuration(-24 * time.Hour)},
		{time.Date(2009, 11, 24, 0, 0, 0, 0, time.UTC), time.Date(2009, 11, 23, 0, 0, 0, 0, time.UTC), NewSimpleDuration(24 * time.Hour)},
		{time.Date(-2009, 11, 24, 0, 0, 0, 0, time.UTC), time.Date(-2009, 11, 23, 0, 0, 0, 0, time.UTC), NewSimpleDuration(24 * time.Hour)},
		{time.Time{}, time.Date(2109, 11, 23, 0, 0, 0, 0, time.UTC), NewSimpleDuration(minDuration)},
		{time.Date(2109, 11, 23, 0, 0, 0, 0, time.UTC), time.Time{}, NewSimpleDuration(maxDuration)},
		{time.Time{}, time.Date(-2109, 11, 23, 0, 0, 0, 0, time.UTC), NewSimpleDuration(maxDuration)},
		{time.Date(-2109, 11, 23, 0, 0, 0, 0, time.UTC), time.Time{}, NewSimpleDuration(minDuration)},
		{time.Date(2290, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), NewSimpleDuration(290*365*24*time.Hour + 71*24*time.Hour)},
		{time.Date(2300, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), NewSimpleDuration(maxDuration)},
		{time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2290, 1, 1, 0, 0, 0, 0, time.UTC), NewSimpleDuration(-290*365*24*time.Hour - 71*24*time.Hour)},
		{time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2300, 1, 1, 0, 0, 0, 0, time.UTC), NewSimpleDuration(minDuration)},
	}
	for i, test := range tests {
		result := NewSimpleTime(test.t).Since(test.u)
		if !result.Compare(test.d) {
			t.Errorf("TestSince %d Failed:Expected=(%v),\nResult=(%v)", i, test.d, result)
		}
	}
}
