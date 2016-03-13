// Package simpletime is an extension of the go standard library time
// package which provides a number of convenience functions for retrieving
// new times based on the provided reference time
package simpletime

import (
	"time"
)

// SimpleTime struct embeds time.Time
type SimpleTime struct {
	time.Time
}

// NewSimpleTime returns a pointer to a SimpleTime struct
func NewSimpleTime(t time.Time) *SimpleTime {
	return &SimpleTime{t}
}

// BeginningOfYear returns the beginning of the year from the SimpleTime struct
func (t *SimpleTime) BeginningOfYear() time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// BeginningOfMonth returns the beginning of the month from the SimpleTime struct
func (t *SimpleTime) BeginningOfMonth() time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// BeginningOfDay returns the beginning of the day from the SimpleTime struct
func (t *SimpleTime) BeginningOfDay() time.Time {
	return t.Truncate(time.Hour * 24)
}

// BeginningOfHour returns the beginning of the hour from the SimpleTime struct
func (t *SimpleTime) BeginningOfHour() time.Time {
	return t.Truncate(time.Hour)
}

// BeginningOfMinute returns the beginning of the minute from the SimpleTime struct
func (t *SimpleTime) BeginningOfMinute() time.Time {
	return t.Truncate(time.Minute)
}

// BeginningOfSecond returns the beginning of the second from the SimpleTime struct
func (t *SimpleTime) BeginningOfSecond() time.Time {
	return t.Truncate(time.Second)
}

// EndOfYear returns the end of the year from the SimpleTime struct
func (t *SimpleTime) EndOfYear() time.Time {
	return t.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// EndOfMonth returns the end of the month from the SimpleTime struct
func (t *SimpleTime) EndOfMonth() time.Time {
	return t.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfDay returns the end of the day from the SimpleTime struct
func (t *SimpleTime) EndOfDay() time.Time {
	return t.BeginningOfDay().AddDate(0, 0, 1).Add(-time.Nanosecond)
}

// EndOfHour returns the end of the hour from the SimpleTime struct
func (t *SimpleTime) EndOfHour() time.Time {
	return t.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

// EndOfMinute returns the end of the minute from the SimpleTime struct
func (t *SimpleTime) EndOfMinute() time.Time {
	return t.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

// EndOfSecond returns the end of the second from the SimpleTime struct
func (t *SimpleTime) EndOfSecond() time.Time {
	return t.BeginningOfSecond().Add(time.Second - time.Nanosecond)
}

// NextOccurrenceOfWeekday returns the next coming or occurence of the
// time.Weekday argument passed in
func (t *SimpleTime) NextOccurrenceOfWeekday(weekday time.Weekday) time.Time {
	days := (int(weekday) - int(t.Weekday()) + 7) % 7
	next := NewSimpleTime(t.AddDate(0, 0, days))
	if days > 0 {
		return next.BeginningOfDay()
	}
	return next.Time
}

// NextSunday returns the next coming Sunday from the
// date of the SimpleTime struct
func (t *SimpleTime) NextSunday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(0))
}

// NextMonday returns the next coming Monday from the
// date of the SimpleTime struct
func (t *SimpleTime) NextMonday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(1))
}

// NextTuesday returns the next coming Tuesday from the
// date of the SimpleTime struct
func (t *SimpleTime) NextTuesday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(2))
}

// NextWednesday returns the next coming Wednesday from the
// date of the SimpleTime struct
func (t *SimpleTime) NextWednesday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(3))
}

// NextThursday returns the next coming Thursday from the
// date of the SimpleTime struct
func (t *SimpleTime) NextThursday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(4))
}

// NextFriday returns the next coming Friday from the
// date of the SimpleTime struct
func (t *SimpleTime) NextFriday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(5))
}

// NextSaturday returns the next coming Saturday from the
// date of the SimpleTime struct
func (t *SimpleTime) NextSaturday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(6))
}

// AddDays adds n days to the SimpleTime date
func (t *SimpleTime) AddDays(days int) time.Time {
	return t.AddDate(0, 0, days)
}

// NextDay returns the next day in time.Time format
func (t *SimpleTime) NextDay() time.Time {
	return t.AddDays(1)
}

// PrevDay subtracts one day
func (t *SimpleTime) PrevDay() time.Time {
	return t.AddDays(-1)
}

// Since returns the time elapsed since u from the reference time t
func (t *SimpleTime) Since(u time.Time) *SimpleDuration {
	d := t.Sub(u)
	return NewSimpleDuration(d)
}
