package timeparse

import (
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	testData := []struct {
		Layout string
		Time   string
		Want   time.Time
	}{
		{
			Layout: "2006 01 02",
			Time:   "2021 12 24",
			Want:   time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, test := range testData {
		got, err := time.Parse(test.Layout, test.Time)
		if err != nil {
			t.Error(err)
		} else {
			if test.Want != got {
				t.Errorf("Parse time=%s, layout=%s, want=%v, got=%v", test.Time, test.Layout, test.Want, got)
			}
		}

	}
}
