package simpletime

import (
	"testing"
	"time"
)

const (
	minDuration time.Duration = -1 << 63
	maxDuration time.Duration = 1<<63 - 1
)

func TestCompare(t *testing.T) {
	tests := []struct {
		d1       *SimpleDuration
		d2       *SimpleDuration
		expected bool
	}{
		{nil, nil, false},
		{NewSimpleDuration(0), nil, false},
		{&SimpleDuration{}, &SimpleDuration{}, true},
		{NewSimpleDuration(0), NewSimpleDuration(0), true},
		{NewSimpleDuration(minDuration), NewSimpleDuration(minDuration), true},
		{NewSimpleDuration(maxDuration), NewSimpleDuration(maxDuration), true},
		{NewSimpleDuration(minDuration), NewSimpleDuration(maxDuration), false},
	}
	for i, test := range tests {
		result := test.d1.Compare(test.d2)
		if result != test.expected {
			t.Errorf("TestCompare %d Failed:Expected=(%v),\nResult=(%v)", i, test.expected, result)
		}
	}
}

func TestDays(t *testing.T) {
	tests := []struct {
		d        *SimpleDuration
		expected float64
	}{
		{NewSimpleDuration(0), 0},
		{NewSimpleDuration(86400 * time.Second), 1.0},
		{NewSimpleDuration(365 * time.Hour * 24), 365},
		{NewSimpleDuration(-365 * time.Hour * 24), -365},
		{NewSimpleDuration(maxDuration), 106751.99116730066},
		{NewSimpleDuration(minDuration), -106751.99116730066},
	}
	for i, test := range tests {
		result := test.d.Days()
		if result != test.expected {
			t.Errorf("TestDays %d Failed:Expected=(%v),\nResult=(%v)", i, test.expected, result)
		}
	}
}
