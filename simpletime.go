package simpletime

import (
	"time"
)

type SimpleTime struct {
	time.Time
}

func NewSimpleTime(t time.Time) *SimpleTime {
	return &SimpleTime{t}
}

func (t *SimpleTime) BeginningOfYear() time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

func (t *SimpleTime) BeginningOfMonth() time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func (t *SimpleTime) BeginningOfDay() time.Time {
	return t.Truncate(time.Hour * 24)
}

func (t *SimpleTime) BeginningOfHour() time.Time {
	return t.Truncate(time.Hour)
}

func (t *SimpleTime) BeginningOfMinute() time.Time {
	return t.Truncate(time.Minute)
}

func (t *SimpleTime) BeginningOfSecond() time.Time {
	return t.Truncate(time.Second)
}

func (t *SimpleTime) EndOfYear() time.Time {
	return t.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func (t *SimpleTime) EndOfMonth() time.Time {
	return t.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (t *SimpleTime) EndOfDay() time.Time {
	return t.BeginningOfDay().AddDate(0, 0, 1).Add(-time.Nanosecond)
}

func (t *SimpleTime) EndOfHour() time.Time {
	return t.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

func (t *SimpleTime) EndOfMinute() time.Time {
	return t.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

func (t *SimpleTime) EndOfSecond() time.Time {
	return t.BeginningOfSecond().Add(time.Second - time.Nanosecond)
}

func (t *SimpleTime) NextOccurrenceOfWeekday(weekday time.Weekday) time.Time {
	days := (int(weekday) - int(t.Weekday()) + 7) % 7
	next := NewSimpleTime(t.AddDate(0, 0, days))
	if days > 0 {
		return next.BeginningOfDay()
	}
	return next.Time
}

func (t *SimpleTime) NextSunday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(0))
}

func (t *SimpleTime) NextMonday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(1))
}

func (t *SimpleTime) NextTuesday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(2))
}

func (t *SimpleTime) NextWednesday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(3))
}

func (t *SimpleTime) NextThursday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(4))
}

func (t *SimpleTime) NextFriday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(5))
}

func (t *SimpleTime) NextSaturday() time.Time {
	return t.NextOccurrenceOfWeekday(time.Weekday(6))
}

func (t *SimpleTime) AddDays(days int) time.Time {
	return t.AddDate(0, 0, days)
}

func (t *SimpleTime) NextDay() time.Time {
	return t.AddDays(1)
}
