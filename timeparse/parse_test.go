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
		{
			Layout: "2006 1 2",
			Time:   "2021 12 24",
			Want:   time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 1 _2",
			Time:   "2021 12  5",
			Want:   time.Date(2021, 12, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 __2",
			Time:   "2021 145",
			Want:   time.Date(2021, 5, 25, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 002",
			Time:   "2021 145",
			Want:   time.Date(2021, 5, 25, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 __2",
			Time:   "2021  31",
			Want:   time.Date(2021, 1, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 002",
			Time:   "2021 031",
			Want:   time.Date(2021, 1, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 January 2",
			Time:   "2021 February 28",
			Want:   time.Date(2021, 2, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 Jan 2",
			Time:   "2021 Feb 28",
			Want:   time.Date(2021, 2, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 01 02 Monday",
			Time:   "2021 10 04 Monday",
			Want:   time.Date(2021, 10, 04, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 01 02 Monday",
			Time:   "2021 10 04 Friday",// for parse this is ignored but has to be valid day
			Want:   time.Date(2021, 10, 04, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 01 02 Mon",
			Time:   "2021 10 04 Mon",
			Want:   time.Date(2021, 10, 04, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "2006 01 02 Mon",
			Time:   "2021 10 04 Fri",// for parse this is ignored but has to be valid day
			Want:   time.Date(2021, 10, 04, 0, 0, 0, 0, time.UTC),
		},
		{
			Layout: "15:04:05",
			Time:   "23:55:55",
			Want:   time.Date(0, 1, 1, 23, 55, 55, 0, time.UTC),
		},
		{
			Layout: "3:04:05 PM",
			Time:   "12:55:55 PM",
			Want:   time.Date(0, 1, 1, 12, 55, 55, 0, time.UTC),
		},
		{
			Layout: "03:4:5 pm 06",
			Time:   "12:55:55 am 20",
			Want:   time.Date(2020, 1, 1, 0, 55, 55, 0, time.UTC),
		},
		{
			Layout: "03:4:5.000000000 pm 06",
			Time:   "12:55:55.123456000 am 20",
			Want:   time.Date(2020, 1, 1, 0, 55, 55, 123456000, time.UTC),
		},
		{
			Layout: "03:4:5.999999999 pm 06",
			Time:   "12:55:55.123456 am 20",
			Want:   time.Date(2020, 1, 1, 0, 55, 55, 123456000, time.UTC),
		},
		{
			Layout: "03:4:5,000000000 pm 06",
			Time:   "12:55:55,123456000 am 20",
			Want:   time.Date(2020, 1, 1, 0, 55, 55, 123456000, time.UTC),
		},
		{
			Layout: "03:4:5,999999999 pm 06",
			Time:   "12:55:55,123456 am 20",
			Want:   time.Date(2020, 1, 1, 0, 55, 55, 123456000, time.UTC),
		},
		{
			Layout: "03:4:5,999999999 pm 06",
			Time:   "12:55:55.123456 am 20",
			Want:   time.Date(2020, 1, 1, 0, 55, 55, 123456000, time.UTC),
		},
		{
			Layout: "03:4:5.999999999 pm 06",
			Time:   "12:55:55,123456 am 20",
			Want:   time.Date(2020, 1, 1, 0, 55, 55, 123456000, time.UTC),
		},
		// {
		// 	Layout: "2006 __ 01 02",
		// 	Time:   "2021 53 12 31",
		// 	Want:   time.Date(2021, 2, 31, 0, 0, 0, 0, time.UTC),
		// },
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

func TestTimeParseZone(t *testing.T) {
	testData := []struct {
		Layout string
		Time   string
		Want   time.Time
	}{
		{
			Layout: "2006 01 02 15:04:05 MST",
			Time:   "2021 01 01 12:55:55 CET",
			Want:   time.Date(2021, 1, 1, 12, 55, 55, 0, location("CET")),
		},
		{
			Layout: "2006 01 02 15:04:05Z070000",
			Time:   "2021 01 01 12:55:55-050000",
			Want:   time.Date(2021, 1, 1, 12, 55, 55, 0, location("EST")),
		},
		{
			Layout: "2006 01 02 15:04:05-07:00:00",
			Time:   "2021 01 01 12:55:55-05:00:00",
			Want:   time.Date(2021, 1, 1, 12, 55, 55, 0, location("EST")),
		},
		// {
		// 	Layout: "2006 01 02 15:04:05 MST",
		// 	Time:   "2021 01 01 12:55:55 Asia/Shanghai", // not possible
		// 	Want:   time.Date(2021, 1, 1, 12, 55, 55, 0, location("EST")),
		// },
	}
	for _, test := range testData {
		got, err := time.Parse(test.Layout, test.Time)
		if err != nil {
			t.Error(err)
		} else {
			if test.Want.UnixMilli() != got.UnixMilli() {
				t.Errorf("Parse time=%s, layout=%s, want=%v, got=%v", test.Time, test.Layout, test.Want, got)
			}
		}

	}
}

func zone(zone string) time.Time {
	location := location(zone)
	return time.Date(2021, 1, 1, 1, 1, 1, 111111111, location)
}

func location(zone string) *time.Location {
	loc, _ := time.LoadLocation(zone)
	return loc
}
