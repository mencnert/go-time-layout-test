package timeformat

import (
	"fmt"
	"log"
	"testing"
	"time"

	timefmt "github.com/itchyny/timefmt-go"
)

func TestGoTimeFormat(t *testing.T) {
	testData := []struct {
		Timestamp      time.Time
		GoLayout       string
		StrftimeLayout string
		Expected       string
	}{
		// Long year
		{
			Timestamp:      time.Date(2021,1,1,0,0,0,0,time.UTC),
			GoLayout:       "2006",
			StrftimeLayout: "%Y",
			Expected:       "2021",
		},
		{
			Timestamp:      time.Date(1981,1,1,0,0,0,0,time.UTC),
			GoLayout:       "2006",
			StrftimeLayout: "%Y",
			Expected:       "1981",
		},
		{
			Timestamp:      time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:       "2006",
			StrftimeLayout: "%Y",
			Expected:       "0001",
		},
		// 3 digit year
		// here first zero is just digit not part of Go time layout
		{
			Timestamp:      time.Date(1981,1,1,0,0,0,0,time.UTC),
			GoLayout:  "006",
			Expected:  "081",
		},
		// 2 digit year
		{
			Timestamp:      time.Date(2021,1,1,0,0,0,0,time.UTC),
			GoLayout:       "06",
			StrftimeLayout: "%y",
			Expected:       "21",
		},
		{
			Timestamp:      time.Date(2001,1,1,0,0,0,0,time.UTC),
			GoLayout:       "06",
			StrftimeLayout: "%y",
			Expected:       "01",
		},
		{
			Timestamp:      time.Date(1981,1,1,0,0,0,0,time.UTC),
			GoLayout:       "06",
			StrftimeLayout: "%y",
			Expected:       "81",
		},
		{
			Timestamp:      time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:       "06",
			StrftimeLayout: "%y",
			Expected:       "01",
		},
		// month long word
		{
			Timestamp: time.Date(1,time.January,1,0,0,0,0,time.UTC),
			GoLayout:  "January",
			Expected:  "January",
		},
		{
			Timestamp: time.Date(1,time.February,1,0,0,0,0,time.UTC),
			GoLayout:  "January",
			Expected:  "February",
		},
		{
			Timestamp: time.Date(1,time.December,1,0,0,0,0,time.UTC),
			GoLayout:  "January",
			Expected:  "December",
		},
		{
			Timestamp: time.Date(1,0,1,0,0,0,0,time.UTC),
			GoLayout:  "January",
			Expected:  "December",
		},
		{
			Timestamp: time.Date(1,13,1,0,0,0,0,time.UTC),
			GoLayout:  "January",
			Expected:  "January",
		},
		// month short word
		{
			Timestamp: time.Date(1,time.January,1,0,0,0,0,time.UTC),
			GoLayout:  "Jan",
			Expected:  "Jan",
		},
		{
			Timestamp: time.Date(1,time.February,1,0,0,0,0,time.UTC),
			GoLayout:  "Jan",
			Expected:  "Feb",
		},
		{
			Timestamp: time.Date(1,time.December,1,0,0,0,0,time.UTC),
			GoLayout:  "Jan",
			Expected:  "Dec",
		},
		// month short number
		{
			Timestamp: time.Date(1,time.January,1,0,0,0,0,time.UTC),
			GoLayout:  "1",
			Expected:  "1",
		},
		{
			Timestamp: time.Date(1,time.February,1,0,0,0,0,time.UTC),
			GoLayout:  "1",
			Expected:  "2",
		},
		{
			Timestamp: time.Date(1,time.December,1,0,0,0,0,time.UTC),
			GoLayout:  "1",
			Expected:  "12",
		},
		// month long number
		{
			Timestamp: time.Date(1,time.January,1,0,0,0,0,time.UTC),
			GoLayout:  "01",
			Expected:  "01",
		},
		{
			Timestamp: time.Date(1,time.February,1,0,0,0,0,time.UTC),
			GoLayout:  "01",
			Expected:  "02",
		},
		{
			Timestamp: time.Date(1,time.December,1,0,0,0,0,time.UTC),
			GoLayout:  "01",
			Expected:  "12",
		},
		// day long word
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "Monday",
			Expected:  "Monday",
		},
		{
			Timestamp: time.Date(1,1,3,0,0,0,0,time.UTC),
			GoLayout:  "Monday",
			Expected:  "Wednesday",
		},
		{
			Timestamp: time.Date(1,1,7,0,0,0,0,time.UTC),
			GoLayout:  "Monday",
			Expected:  "Sunday",
		},
		{
			Timestamp: time.Date(1,1,8,0,0,0,0,time.UTC),
			GoLayout:  "Monday",
			Expected:  "Monday",
		},
		// day short word
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "Mon",
			Expected:  "Mon",
		},
		{
			Timestamp: time.Date(1,1,3,0,0,0,0,time.UTC),
			GoLayout:  "Mon",
			Expected:  "Wed",
		},
		{
			Timestamp: time.Date(1,1,7,0,0,0,0,time.UTC),
			GoLayout:  "Mon",
			Expected:  "Sun",
		},
		// day short number
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "2",
			Expected:  "1",
		},
		{
			Timestamp: time.Date(1,1,11,0,0,0,0,time.UTC),
			GoLayout:  "2",
			Expected:  "11",
		},
		{
			Timestamp: time.Date(1,1,31,0,0,0,0,time.UTC),
			GoLayout:  "2",
			Expected:  "31",
		},
		// day zero prefix number
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "02",
			Expected:  "01",
		},
		{
			Timestamp: time.Date(1,1,31,0,0,0,0,time.UTC),
			GoLayout:  "02",
			Expected:  "31",
		},
		// day two zero prefix number
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "002",
			Expected:  "001",
		},
		{
			Timestamp: time.Date(1,1,31,0,0,0,0,time.UTC),
			GoLayout:  "002",
			Expected:  "031",
		},
		{
			Timestamp: time.Date(1,5,25,0,0,0,0,time.UTC),
			GoLayout:  "002",
			Expected:  "145",
		},
		{
			Timestamp: time.Date(1,12,31,0,0,0,0,time.UTC),
			GoLayout:  "002",
			Expected:  "365",
		},
		{
			Timestamp: time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:  "002",
			Expected:  "366",
		},
		// day space prefix if one digit
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "_2",
			Expected:  " 1",
		},
		{
			Timestamp: time.Date(1,1,31,0,0,0,0,time.UTC),
			GoLayout:  "_2",
			Expected:  "31",
		},
		// day space prefix if two digits and two spaces if one digit
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "__2",
			Expected:  "  1",
		},
		{
			Timestamp: time.Date(1,1,31,0,0,0,0,time.UTC),
			GoLayout:  "__2",
			Expected:  " 31",
		},
		{
			Timestamp: time.Date(1,5,25,0,0,0,0,time.UTC),
			GoLayout:  "__2",
			Expected:  "145",
		},
		{
			Timestamp: time.Date(1,12,31,0,0,0,0,time.UTC),
			GoLayout:  "__2",
			Expected:  "365",
		},
		{
			Timestamp: time.Date(2020,12,31,0,0,0,0,time.UTC),
			GoLayout:  "__2",
			Expected:  "366",
		},
		// hour 24h format
		{
			Timestamp: time.Date(1,1,1,1,0,0,0,time.UTC),
			GoLayout:  "15",
			Expected:  "01",
		},
		{
			Timestamp: time.Date(1,1,1,12,0,0,0,time.UTC),
			GoLayout:  "15",
			Expected:  "12",
		},
		{
			Timestamp: time.Date(1,1,1,15,0,0,0,time.UTC),
			GoLayout:  "15",
			Expected:  "15",
		},
		{
			Timestamp: time.Date(1,1,1,23,0,0,0,time.UTC),
			GoLayout:  "15",
			Expected:  "23",
		},
		{
			Timestamp: time.Date(1,1,1,24,0,0,0,time.UTC),
			GoLayout:  "15",
			Expected:  "00",
		},
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "15",
			Expected:  "00",
		},
		// hour 12 hour system short
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "3 PM",
			Expected:  "12 AM",
		},
		{
			Timestamp: time.Date(1,1,1,1,0,0,0,time.UTC),
			GoLayout:  "3 PM",
			Expected:  "1 AM",
		},
		{
			Timestamp: time.Date(1,1,1,12,0,0,0,time.UTC),
			GoLayout:  "3 PM",
			Expected:  "12 PM",
		},
		{
			Timestamp: time.Date(1,1,1,15,0,0,0,time.UTC),
			GoLayout:  "3 PM",
			Expected:  "3 PM",
		},
		{
			Timestamp: time.Date(1,1,1,24,0,0,0,time.UTC),
			GoLayout:  "3 PM",
			Expected:  "12 AM",
		},
		// hour 12 hour system short
		{
			Timestamp: time.Date(1,1,1,0,0,0,0,time.UTC),
			GoLayout:  "03 PM",
			Expected:  "12 AM",
		},
		{
			Timestamp: time.Date(1,1,1,1,0,0,0,time.UTC),
			GoLayout:  "03 PM",
			Expected:  "01 AM",
		},
		{
			Timestamp: time.Date(1,1,1,12,0,0,0,time.UTC),
			GoLayout:  "03 PM",
			Expected:  "12 PM",
		},
		{
			Timestamp: time.Date(1,1,1,15,0,0,0,time.UTC),
			GoLayout:  "03 PM",
			Expected:  "03 PM",
		},
		{
			Timestamp: time.Date(1,1,1,24,0,0,0,time.UTC),
			GoLayout:  "03 PM",
			Expected:  "12 AM",
		},
		// minute short
		{
			Timestamp: minute(0),
			GoLayout:  "4",
			Expected:  "0",
		},
		{
			Timestamp: minute(4),
			GoLayout:  "4",
			Expected:  "4",
		},
		{
			Timestamp: minute(10),
			GoLayout:  "4",
			Expected:  "10",
		},
		{
			Timestamp: minute(60),
			GoLayout:  "4",
			Expected:  "0",
		},
		// minute long
		{
			Timestamp: minute(0),
			GoLayout:  "04",
			Expected:  "00",
		},
		{
			Timestamp: minute(4),
			GoLayout:  "04",
			Expected:  "04",
		},
		{
			Timestamp: minute(10),
			GoLayout:  "04",
			Expected:  "10",
		},
		{
			Timestamp: minute(60),
			GoLayout:  "04",
			Expected:  "00",
		},
		// second short
		{
			Timestamp: second(0),
			GoLayout:  "5",
			Expected:  "0",
		},
		{
			Timestamp: second(5),
			GoLayout:  "5",
			Expected:  "5",
		},
		{
			Timestamp: second(25),
			GoLayout:  "5",
			Expected:  "25",
		},
		{
			Timestamp: second(60),
			GoLayout:  "5",
			Expected:  "0",
		},
		// second long
		{
			Timestamp: second(0),
			GoLayout:  "05",
			Expected:  "00",
		},
		{
			Timestamp: second(5),
			GoLayout:  "05",
			Expected:  "05",
		},
		{
			Timestamp: second(25),
			GoLayout:  "05",
			Expected:  "25",
		},
		// part of day upper case
		{
			Timestamp: hour(0),
			GoLayout:  "PM",
			Expected:  "AM",
		},
		{
			Timestamp: hour(5),
			GoLayout:  "PM",
			Expected:  "AM",
		},
		{
			Timestamp: hour(12),
			GoLayout:  "PM",
			Expected:  "PM",
		},
		{
			Timestamp: hour(15),
			GoLayout:  "PM",
			Expected:  "PM",
		},
		{
			Timestamp: hour(23),
			GoLayout:  "PM",
			Expected:  "PM",
		},
		{
			Timestamp: hour(24),
			GoLayout:  "PM",
			Expected:  "AM",
		},
		// part of day lower case
		{
			Timestamp: hour(0),
			GoLayout:  "pm",
			Expected:  "am",
		},
		{
			Timestamp: hour(5),
			GoLayout:  "pm",
			Expected:  "am",
		},
		{
			Timestamp: hour(12),
			GoLayout:  "pm",
			Expected:  "pm",
		},
		{
			Timestamp: hour(15),
			GoLayout:  "pm",
			Expected:  "pm",
		},
		{
			Timestamp: hour(23),
			GoLayout:  "pm",
			Expected:  "pm",
		},
		{
			Timestamp: hour(24),
			GoLayout:  "pm",
			Expected:  "am",
		},
		// milliseconds zero omited
		{
			Timestamp: millisecond(123),
			GoLayout:  "000", // no dot
			Expected:  "000",
		},
		{
			Timestamp: millisecond(123),
			GoLayout:  ".0",
			Expected:  ".1",
		},
		{
			Timestamp: millisecond(123),
			GoLayout:  ".00",
			Expected:  ".12",
		},
		{
			Timestamp: millisecond(123),
			GoLayout:  ".000",
			Expected:  ".123",
		},
		{
			Timestamp: millisecond(199),
			GoLayout:  ".00",
			Expected:  ".19",
		},
		{
			Timestamp: millisecond(199),
			GoLayout:  ".000000",
			Expected:  ".199000",
		},
		{
			Timestamp: millisecond(199),
			GoLayout:  ".000000000",
			Expected:  ".199000000",
		},
		{
			Timestamp: millisecond(199),
			GoLayout:  ",000000",
			Expected:  ".199000",
		},
		// milliseconds trailing zeros omited
		{
			Timestamp: millisecond(199),
			GoLayout:  "99", // no dot
			Expected:  "99",
		},
		{
			Timestamp: millisecond(199),
			GoLayout:  ".999",
			Expected:  ".199",
		},
		{
			Timestamp: millisecond(199),
			GoLayout:  ".999999",
			Expected:  ".199",
		},
		{
			Timestamp: millisecond(199),
			GoLayout:  "01,999999",
			Expected:  "01.199",
		},
		{
			Timestamp: millisecond(199),
			GoLayout:  ".90000", // not possible to combine
			Expected:  ".90000",
		},
		{
			Timestamp: millisecond(199),
			GoLayout:  ".9999 .0000",
			Expected:  ".199 .1990",
		},
		//time zone
		{
			Timestamp: zone("UTC"),
			GoLayout:  "MST",
			Expected:  "UTC",
		},
		{
			Timestamp: zone("CET"),
			GoLayout:  "MST",
			Expected:  "CET",
		},
		//time zone Z0700
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z0700",
			Expected:  "Z",
		},
		{
			Timestamp: zone("CET"),
			GoLayout:  "Z0700",
			Expected:  "+0100",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z0700",
			Expected:  "+0800",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z0700",
			Expected:  "-0500",
		},
		//time zone Z070000
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z070000",
			Expected:  "Z",
		},
		{
			Timestamp: zone("CET"),
			GoLayout:  "Z070000",
			Expected:  "+010000",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z070000",
			Expected:  "+080000",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z070000",
			Expected:  "-050000",
		},
		//time zone Z07
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z07",
			Expected:  "Z",
		},
		{
			Timestamp: zone("CET"),
			GoLayout:  "Z07",
			Expected:  "+01",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z07",
			Expected:  "+08",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z07",
			Expected:  "-05",
		},
		//time zone Z07:00
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z07:00",
			Expected:  "Z",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z07:00",
			Expected:  "+08:00",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z07:00",
			Expected:  "-05:00",
		},
		//time zone Z07:00:00
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z07:00:00",
			Expected:  "Z",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z07:00:00",
			Expected:  "+08:00:00",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z07:00:00",
			Expected:  "-05:00:00",
		},
		//time zone -07
		{
			Timestamp: zone("UTC"),
			GoLayout:  "-07",
			Expected:  "+00",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "-07",
			Expected:  "+08",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "-07",
			Expected:  "-05",
		},
		//time zone -0700
		{
			Timestamp: zone("UTC"),
			GoLayout:  "-0700",
			Expected:  "+0000",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "-0700",
			Expected:  "+0800",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "-0700",
			Expected:  "-0500",
		},
		//time zone -070000
		{
			Timestamp: zone("UTC"),
			GoLayout:  "-070000",
			Expected:  "+000000",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "-070000",
			Expected:  "+080000",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "-070000",
			Expected:  "-050000",
		},
		//time zone -07:00
		{
			Timestamp: zone("UTC"),
			GoLayout:  "-07:00",
			Expected:  "+00:00",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "-07:00",
			Expected:  "+08:00",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "-07:00",
			Expected:  "-05:00",
		},
		//time zone -07:00:00
		{
			Timestamp: zone("UTC"),
			GoLayout:  "-07:00:00",
			Expected:  "+00:00:00",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "-07:00:00",
			Expected:  "+08:00:00",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "-07:00:00",
			Expected:  "-05:00:00",
		},
		// complex
		{
			Timestamp: time.Date(2021, 2, 20, 23, 22, 21, 123456, location("Asia/Shanghai")),
			GoLayout:  "January Jan 1 01 Monday Mon 2 02 002 _2 __2 15 3 03 4 04 5 05 06 2006 PM pm .000000000 .999999999 MST Z07 Z0700 Z070000 Z07:00 Z07:00:00 -07 -0700 -070000 -07:00 -07:00:00",
			Expected:  "February Feb 2 02 Saturday Sat 20 20 051 20  51 23 11 11 22 22 21 21 21 2021 PM pm .000123456 .000123456 CST +08 +0800 +080000 +08:00 +08:00:00 +08 +0800 +080000 +08:00 +08:00:00",
		},
		// SOF
		{
			Timestamp:      time.Date(2021, 2, 12, 15, 05, 03, 123, time.UTC),
			GoLayout:       "20060102150405",
			StrftimeLayout: "%Y%m%d%H%M%S",
			Expected:       "20210212150503",
		},
	}

	for _, test := range testData {
		actualGoResult := test.Timestamp.Format(test.GoLayout)
		actualStrftimeResult := timefmt.Format(test.Timestamp, test.StrftimeLayout)

		if test.Expected != actualGoResult {
			t.Errorf("\n%-10s %v\n%-10s %s\n%-10s %s\n%-10s %s",
				"time", test.Timestamp,
				"template", test.GoLayout,
				"got", actualGoResult,
				"want", test.Expected,
			)
		}

		if test.StrftimeLayout != "" {
			if actualStrftimeResult != test.Expected {
				t.Errorf("\n%-10s %v\n%-10s %s\n%-10s %s\n%-10s %s",
					"time", test.Timestamp,
					"template", test.StrftimeLayout,
					"got", actualStrftimeResult,
					"want", test.Expected,
				)
			}
		}
	}
}

func TestIsoWeek(t *testing.T) {
	testData := []struct {
		timestamp time.Time
		expected  int
	}{
		{
			expected:  1,
			timestamp: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			expected:  53,
			timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			expected:  25,
			timestamp: time.Date(2021, 6, 27, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range testData {
		_, week := test.timestamp.ISOWeek()
		if test.expected != week {
			t.Errorf("timestamp %v\nwant %d\ngot %d", test.timestamp, test.expected, week)
		}
	}
}

func TestStrfTime(t *testing.T) {
	tt, err := timefmt.Parse("2020/07/24 09:07:29", "%Y/%m/%d %H:%M:%S")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tt) // 2020-07-24 09:07:29 +0000 UTC

	str := timefmt.Format(tt, "%Y/%m/%d %H:%M:%S")
	fmt.Println(str) // 2020/07/24 09:07:29

	str = timefmt.Format(tt, "%a, %d %b %Y %T %z")
	fmt.Println(str) // Fri, 24 Jul 2020 09:07:29 +0000
}

func year(y int) time.Time {
	return time.Date(y, 1, 1, 1, 1, 1, 1, time.UTC)
}

func month(m time.Month) time.Time {
	return time.Date(1, m, 1, 1, 1, 1, 1, time.UTC)
}

func day(d int) time.Time {
	return time.Date(1, 1, d, 1, 1, 1, 1, time.UTC)
}

func hour(h int) time.Time {
	return time.Date(1, 1, 1, h, 0, 0, 0, time.UTC)
}

func minute(m int) time.Time {
	return time.Date(1, 1, 1, 1, m, 0, 0, time.UTC)
}

func second(s int) time.Time {
	return time.Date(1, 1, 1, 1, 1, s, 0, time.UTC)
}

func millisecond(ms int) time.Time {
	ms = ms * int(time.Millisecond)
	return time.Date(1, 1, 1, 1, 1, 1, ms, time.UTC)
}

func zone(zone string) time.Time {
	location, _ := time.LoadLocation(zone)
	return time.Date(2021, 1, 1, 1, 1, 1, 111111111, location)
}

func monthDay(m time.Month, d int) time.Time {
	return time.Date(1, m, d, 1, 1, 1, 1, time.UTC)
}

func location(zone string) *time.Location {
	loc, _ := time.LoadLocation(zone)
	return loc
}
