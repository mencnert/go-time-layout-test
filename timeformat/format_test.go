package timeformat

import (
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
			Timestamp:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "2006",
			StrftimeLayout: "%Y",
			Expected:       "2021",
		},
		{
			Timestamp:      time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "2006",
			StrftimeLayout: "%Y",
			Expected:       "1981",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "2006",
			StrftimeLayout: "%Y",
			Expected:       "0001",
		},
		// 3 digit year not supported
		// here first zero is just digit not part of Go time layout
		{
			Timestamp:      time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "006",
			StrftimeLayout: "0%y",
			Expected:       "081",
		},
		// 2 digit year
		{
			Timestamp:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "06",
			StrftimeLayout: "%y",
			Expected:       "21",
		},
		{
			Timestamp:      time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "06",
			StrftimeLayout: "%y",
			Expected:       "01",
		},
		{
			Timestamp:      time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "06",
			StrftimeLayout: "%y",
			Expected:       "81",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "06",
			StrftimeLayout: "%y",
			Expected:       "01",
		},
		// month long word
		{
			Timestamp:      time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "January",
			StrftimeLayout: "%B",
			Expected:       "January",
		},
		{
			Timestamp:      time.Date(1, time.February, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "January",
			StrftimeLayout: "%B",
			Expected:       "February",
		},
		{
			Timestamp:      time.Date(1, time.December, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "January",
			StrftimeLayout: "%B",
			Expected:       "December",
		},
		{
			Timestamp:      time.Date(1, 0, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "January",
			StrftimeLayout: "%B",
			Expected:       "December",
		},
		{
			Timestamp:      time.Date(1, 13, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "January",
			StrftimeLayout: "%B",
			Expected:       "January",
		},
		// month short word
		{
			Timestamp:      time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Jan",
			StrftimeLayout: "%b",
			Expected:       "Jan",
		},
		{
			Timestamp:      time.Date(1, time.February, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Jan",
			StrftimeLayout: "%b",
			Expected:       "Feb",
		},
		{
			Timestamp:      time.Date(1, time.December, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Jan",
			StrftimeLayout: "%b",
			Expected:       "Dec",
		},
		// month short number
		{
			Timestamp: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:  "1",
			// StrftimeLayout: "", // N/A
			Expected: "1",
		},
		{
			Timestamp: time.Date(1, time.February, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:  "1",
			// StrftimeLayout: "", // N/A
			Expected: "2",
		},
		{
			Timestamp: time.Date(1, time.December, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:  "1",
			// StrftimeLayout: "", // N/A
			Expected: "12",
		},
		// month long number
		{
			Timestamp:      time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "01",
			StrftimeLayout: "%m",
			Expected:       "01",
		},
		{
			Timestamp:      time.Date(1, time.February, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "01",
			StrftimeLayout: "%m",
			Expected:       "02",
		},
		{
			Timestamp:      time.Date(1, time.December, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "01",
			StrftimeLayout: "%m",
			Expected:       "12",
		},
		// day long word
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Monday",
			StrftimeLayout: "%A",
			Expected:       "Monday",
		},
		{
			Timestamp:      time.Date(1, 1, 3, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Monday",
			StrftimeLayout: "%A",
			Expected:       "Wednesday",
		},
		{
			Timestamp:      time.Date(1, 1, 7, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Monday",
			StrftimeLayout: "%A",
			Expected:       "Sunday",
		},
		{
			Timestamp:      time.Date(1, 1, 8, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Monday",
			StrftimeLayout: "%A",
			Expected:       "Monday",
		},
		// day short word
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Mon",
			StrftimeLayout: "%a",
			Expected:       "Mon",
		},
		{
			Timestamp:      time.Date(1, 1, 3, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Mon",
			StrftimeLayout: "%a",
			Expected:       "Wed",
		},
		{
			Timestamp:      time.Date(1, 1, 7, 0, 0, 0, 0, time.UTC),
			GoLayout:       "Mon",
			StrftimeLayout: "%a",
			Expected:       "Sun",
		},
		// day short number
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:  "2",
			// StrftimeLayout: "", // N/A
			Expected: "1",
		},
		{
			Timestamp: time.Date(1, 1, 11, 0, 0, 0, 0, time.UTC),
			GoLayout:  "2",
			// StrftimeLayout: "", // N/A
			Expected: "11",
		},
		{
			Timestamp: time.Date(1, 1, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:  "2",
			// StrftimeLayout: "", // N/A
			Expected: "31",
		},
		// day zero prefix number
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "02",
			StrftimeLayout: "%d",
			Expected:       "01",
		},
		{
			Timestamp:      time.Date(1, 1, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:       "02",
			StrftimeLayout: "%d",
			Expected:       "31",
		},
		// day of the year
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "002",
			StrftimeLayout: "%j",
			Expected:       "001",
		},
		{
			Timestamp:      time.Date(1, 1, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:       "002",
			StrftimeLayout: "%j",
			Expected:       "031",
		},
		{
			Timestamp:      time.Date(1, 5, 25, 0, 0, 0, 0, time.UTC),
			GoLayout:       "002",
			StrftimeLayout: "%j",
			Expected:       "145",
		},
		{
			Timestamp:      time.Date(1, 12, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:       "002",
			StrftimeLayout: "%j",
			Expected:       "365",
		},
		{
			Timestamp:      time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:       "002",
			StrftimeLayout: "%j",
			Expected:       "366",
		},
		// day space prefix if one digit
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "_2",
			StrftimeLayout: "%e",
			Expected:       " 1",
		},
		{
			Timestamp:      time.Date(1, 1, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:       "_2",
			StrftimeLayout: "%e",
			Expected:       "31",
		},
		// The day of the year space prefix if two digits and two spaces if one digit
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:  "__2",
			// StrftimeLayout: "", // N/A
			Expected: "  1",
		},
		{
			Timestamp: time.Date(1, 1, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:  "__2",
			// StrftimeLayout: "", // N/A
			Expected: " 31",
		},
		{
			Timestamp: time.Date(1, 5, 25, 0, 0, 0, 0, time.UTC),
			GoLayout:  "__2",
			// StrftimeLayout: "", // N/A
			Expected: "145",
		},
		{
			Timestamp: time.Date(1, 12, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:  "__2",
			// StrftimeLayout: "", // N/A
			Expected: "365",
		},
		{
			Timestamp: time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			GoLayout:  "__2",
			// StrftimeLayout: "", // N/A
			Expected: "366",
		},
		// hour 24h format
		{
			Timestamp:      time.Date(1, 1, 1, 1, 0, 0, 0, time.UTC),
			GoLayout:       "15",
			StrftimeLayout: "%H",
			Expected:       "01",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC),
			GoLayout:       "15",
			StrftimeLayout: "%H",
			Expected:       "12",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 15, 0, 0, 0, time.UTC),
			GoLayout:       "15",
			StrftimeLayout: "%H",
			Expected:       "15",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 23, 0, 0, 0, time.UTC),
			GoLayout:       "15",
			StrftimeLayout: "%H",
			Expected:       "23",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 24, 0, 0, 0, time.UTC),
			GoLayout:       "15",
			StrftimeLayout: "%H",
			Expected:       "00",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "15",
			StrftimeLayout: "%H",
			Expected:       "00",
		},
		// hour 12 hour system short
		{
			Timestamp: time.Date(1, 1, 2, 0, 0, 0, 0, time.UTC),
			GoLayout:  "3 PM",
			// StrftimeLayout: "" // N/A %l is space prefixed
			Expected: "12 AM",
		},
		{
			Timestamp: time.Date(1, 1, 1, 1, 0, 0, 0, time.UTC),
			GoLayout:  "3 PM",
			// StrftimeLayout: "" // N/A %l is space prefixed
			Expected: "1 AM",
		},
		{
			Timestamp: time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC),
			GoLayout:  "3 PM",
			// StrftimeLayout: "" // N/A %l is space prefixed
			Expected: "12 PM",
		},
		{
			Timestamp: time.Date(1, 1, 1, 15, 0, 0, 0, time.UTC),
			GoLayout:  "3 PM",
			// StrftimeLayout: "" // N/A %l is space prefixed
			Expected: "3 PM",
		},
		{
			Timestamp: time.Date(1, 1, 1, 24, 0, 0, 0, time.UTC),
			GoLayout:  "3 PM",
			// StrftimeLayout: "" // N/A %l is space prefixed
			Expected: "12 AM",
		},
		// hour 12 hour system short
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "03 PM",
			StrftimeLayout: "%I %p",
			Expected:       "12 AM",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 1, 0, 0, 0, time.UTC),
			GoLayout:       "03 PM",
			StrftimeLayout: "%I %p",
			Expected:       "01 AM",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC),
			GoLayout:       "03 PM",
			StrftimeLayout: "%I %p",
			Expected:       "12 PM",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 15, 0, 0, 0, time.UTC),
			GoLayout:       "03 PM",
			StrftimeLayout: "%I %p",
			Expected:       "03 PM",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 24, 0, 0, 0, time.UTC),
			GoLayout:       "03 PM",
			StrftimeLayout: "%I %p",
			Expected:       "12 AM",
		},
		// minute short
		{
			Timestamp: time.Date(1, 1, 1, 1, 0, 0, 0, time.UTC),
			GoLayout:  "4",
			// StrftimeLayout: "", // N/A
			Expected: "0",
		},
		{
			Timestamp: time.Date(1, 1, 1, 1, 4, 0, 0, time.UTC),
			GoLayout:  "4",
			// StrftimeLayout: "", // N/A
			Expected: "4",
		},
		{
			Timestamp: time.Date(1, 1, 1, 1, 10, 0, 0, time.UTC),
			GoLayout:  "4",
			// StrftimeLayout: "", // N/A
			Expected: "10",
		},
		{
			Timestamp: time.Date(1, 1, 1, 1, 60, 0, 0, time.UTC),
			GoLayout:  "4",
			// StrftimeLayout: "", // N/A
			Expected: "0",
		},
		// minute long
		{
			Timestamp:      time.Date(1, 1, 1, 1, 0, 0, 0, time.UTC),
			GoLayout:       "04",
			StrftimeLayout: "%M",
			Expected:       "00",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 1, 4, 0, 0, time.UTC),
			GoLayout:       "04",
			StrftimeLayout: "%M",
			Expected:       "04",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 1, 10, 0, 0, time.UTC),
			GoLayout:       "04",
			StrftimeLayout: "%M",
			Expected:       "10",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 1, 60, 0, 0, time.UTC),
			GoLayout:       "04",
			StrftimeLayout: "%M",
			Expected:       "00",
		},
		// second short
		{
			Timestamp: time.Date(1, 1, 1, 1, 0, 0, 0, time.UTC),
			GoLayout:  "5",
			// StrftimeLayout: "", // N/A
			Expected: "0",
		},
		{
			Timestamp: time.Date(1, 1, 1, 1, 0, 5, 0, time.UTC),
			GoLayout:  "5",
			// StrftimeLayout: "", // N/A
			Expected: "5",
		},
		{
			Timestamp: time.Date(1, 1, 1, 1, 0, 25, 0, time.UTC),
			GoLayout:  "5",
			// StrftimeLayout: "", // N/A
			Expected: "25",
		},
		{
			Timestamp: time.Date(1, 1, 1, 1, 0, 60, 0, time.UTC),
			GoLayout:  "5",
			// StrftimeLayout: "", // N/A
			Expected: "0",
		},
		// second long
		{
			Timestamp:      time.Date(1, 1, 1, 1, 0, 0, 0, time.UTC),
			GoLayout:       "05",
			StrftimeLayout: "%S",
			Expected:       "00",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 1, 0, 5, 0, time.UTC),
			GoLayout:       "05",
			StrftimeLayout: "%S",
			Expected:       "05",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 1, 0, 25, 0, time.UTC),
			GoLayout:       "05",
			StrftimeLayout: "%S",
			Expected:       "25",
		},
		// part of day upper case
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "PM",
			StrftimeLayout: "%p",
			Expected:       "AM",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 5, 0, 0, 0, time.UTC),
			GoLayout:       "PM",
			StrftimeLayout: "%p",
			Expected:       "AM",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC),
			GoLayout:       "PM",
			StrftimeLayout: "%p",
			Expected:       "PM",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 15, 0, 0, 0, time.UTC),
			GoLayout:       "PM",
			StrftimeLayout: "%p",
			Expected:       "PM",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 23, 0, 0, 0, time.UTC),
			GoLayout:       "PM",
			StrftimeLayout: "%p",
			Expected:       "PM",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 24, 0, 0, 0, time.UTC),
			GoLayout:       "PM",
			StrftimeLayout: "%p",
			Expected:       "AM",
		},
		// part of day lower case
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			GoLayout:       "pm",
			StrftimeLayout: "%P",
			Expected:       "am",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 5, 0, 0, 0, time.UTC),
			GoLayout:       "pm",
			StrftimeLayout: "%P",
			Expected:       "am",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC),
			GoLayout:       "pm",
			StrftimeLayout: "%P",
			Expected:       "pm",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 15, 0, 0, 0, time.UTC),
			GoLayout:       "pm",
			StrftimeLayout: "%P",
			Expected:       "pm",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 23, 0, 0, 0, time.UTC),
			GoLayout:       "pm",
			StrftimeLayout: "%P",
			Expected:       "pm",
		},
		{
			Timestamp:      time.Date(1, 1, 1, 24, 0, 0, 0, time.UTC),
			GoLayout:       "pm",
			StrftimeLayout: "%P",
			Expected:       "am",
		},
		// milliseconds zero omited
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 123000000, time.UTC),
			GoLayout:       "000", // no dot
			StrftimeLayout: "000",
			Expected:       "000",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 123000000, time.UTC),
			GoLayout:  ".0",
			// StrftimeLayout: "", // ??
			Expected: ".1",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 123000000, time.UTC),
			GoLayout:  ".00",
			// StrftimeLayout: "", // ??
			Expected: ".12",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 123000000, time.UTC),
			GoLayout:  ".000",
			// StrftimeLayout: "", // ??
			Expected: ".123",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:  ".00",
			// StrftimeLayout: "", // ??
			Expected: ".19",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:  ".000000",
			// StrftimeLayout: "", // ??
			Expected: ".199000",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:  ".000000000",
			// StrftimeLayout: "", // ??
			Expected: ".199000000",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:  ",000000",
			// StrftimeLayout: "", // ??
			Expected: ".199000",
		},
		// milliseconds trailing zeros omited
		{
			Timestamp:      time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:       "99", // no dot
			StrftimeLayout: "99",
			Expected:       "99",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:  ".999",
			// StrftimeLayout: "", // ??
			Expected: ".199",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:  ".999999",
			// StrftimeLayout: "", // ??
			Expected: ".199",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:  "01,999999",
			// StrftimeLayout: "", // ??
			Expected: "01.199",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:  ".90000", // not possible to combine
			// StrftimeLayout: "", // ??
			Expected: ".90000",
		},
		{
			Timestamp: time.Date(1, 1, 1, 0, 0, 0, 199000000, time.UTC),
			GoLayout:  ".9999 .0000",
			// StrftimeLayout: "", // ??
			Expected: ".199 .1990",
		},
		//time zone
		{
			Timestamp:      zone("UTC"),
			GoLayout:       "MST",
			StrftimeLayout: "%Z",
			Expected:       "UTC",
		},
		{
			Timestamp:      zone("CET"),
			GoLayout:       "MST",
			StrftimeLayout: "%Z",
			Expected:       "CET",
		},
		{
			Timestamp:      zone("Asia/Shanghai"),
			GoLayout:       "MST",
			StrftimeLayout: "%Z",
			Expected:       "CST",
		},
		//time zone Z0700
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z0700",
			// StrftimeLayout: "", // N/A
			Expected: "Z",
		},
		{
			Timestamp: zone("CET"),
			GoLayout:  "Z0700",
			// StrftimeLayout: "", // N/A
			Expected: "+0100",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z0700",
			// StrftimeLayout: "", // N/A
			Expected: "+0800",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z0700",
			// StrftimeLayout: "", // N/A
			Expected: "-0500",
		},
		//time zone Z070000
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z070000",
			// StrftimeLayout: "", // N/A
			Expected: "Z",
		},
		{
			Timestamp: zone("CET"),
			GoLayout:  "Z070000",
			// StrftimeLayout: "", // N/A
			Expected: "+010000",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z070000",
			// StrftimeLayout: "", // N/A
			Expected: "+080000",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z070000",
			// StrftimeLayout: "", // N/A
			Expected: "-050000",
		},
		//time zone Z07
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z07",
			// StrftimeLayout: "", // N/A
			Expected: "Z",
		},
		{
			Timestamp: zone("CET"),
			GoLayout:  "Z07",
			// StrftimeLayout: "", // N/A
			Expected: "+01",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z07",
			// StrftimeLayout: "", // N/A
			Expected: "+08",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z07",
			// StrftimeLayout: "", // N/A
			Expected: "-05",
		},
		//time zone Z07:00
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z07:00",
			// StrftimeLayout: "", // N/A
			Expected: "Z",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z07:00",
			// StrftimeLayout: "", // N/A
			Expected: "+08:00",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z07:00",
			// StrftimeLayout: "", // N/A
			Expected: "-05:00",
		},
		//time zone Z07:00:00
		{
			Timestamp: zone("UTC"),
			GoLayout:  "Z07:00:00",
			// StrftimeLayout: "", // N/A
			Expected: "Z",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "Z07:00:00",
			// StrftimeLayout: "", // N/A
			Expected: "+08:00:00",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "Z07:00:00",
			// StrftimeLayout: "", // N/A
			Expected: "-05:00:00",
		},
		//time zone -07
		{
			Timestamp: zone("UTC"),
			GoLayout:  "-07",
			// StrftimeLayout: "", // N/A
			Expected: "+00",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "-07",
			// StrftimeLayout: "", // N/A
			Expected: "+08",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "-07",
			// StrftimeLayout: "", // N/A
			Expected: "-05",
		},
		//time zone -0700
		{
			Timestamp:      zone("UTC"),
			GoLayout:       "-0700",
			StrftimeLayout: "%z",
			Expected:       "+0000",
		},
		{
			Timestamp:      zone("Asia/Shanghai"),
			GoLayout:       "-0700",
			StrftimeLayout: "%z",
			Expected:       "+0800",
		},
		{
			Timestamp:      zone("America/New_York"),
			GoLayout:       "-0700",
			StrftimeLayout: "%z",
			Expected:       "-0500",
		},
		//time zone -070000
		{
			Timestamp: zone("UTC"),
			GoLayout:  "-070000",
			// StrftimeLayout: "", // N/A
			Expected: "+000000",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "-070000",
			// StrftimeLayout: "", // N/A
			Expected: "+080000",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "-070000",
			// StrftimeLayout: "", // N/A
			Expected: "-050000",
		},
		//time zone -07:00
		{
			Timestamp: zone("UTC"),
			GoLayout:  "-07:00",
			// StrftimeLayout: "", // N/A
			Expected: "+00:00",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "-07:00",
			// StrftimeLayout: "", // N/A
			Expected: "+08:00",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "-07:00",
			// StrftimeLayout: "", // N/A
			Expected: "-05:00",
		},
		//time zone -07:00:00
		{
			Timestamp: zone("UTC"),
			GoLayout:  "-07:00:00",
			// StrftimeLayout: "", // N/A
			Expected: "+00:00:00",
		},
		{
			Timestamp: zone("Asia/Shanghai"),
			GoLayout:  "-07:00:00",
			// StrftimeLayout: "", // N/A
			Expected: "+08:00:00",
		},
		{
			Timestamp: zone("America/New_York"),
			GoLayout:  "-07:00:00",
			// StrftimeLayout: "", // N/A
			Expected: "-05:00:00",
		},
		// complex
		{
			Timestamp: time.Date(2021, 2, 20, 23, 22, 21, 123456, location("Asia/Shanghai")),
			GoLayout:  "January Jan 1 01 Monday Mon 2 02 002 _2 __2 15 3 03 4 04 5 05 06 2006 PM pm .000000000 .999999999 MST Z07 Z0700 Z070000 Z07:00 Z07:00:00 -07 -0700 -070000 -07:00 -07:00:00",
			Expected:  "February Feb 2 02 Saturday Sat 20 20 051 20  51 23 11 11 22 22 21 21 21 2021 PM pm .000123456 .000123456 CST +08 +0800 +080000 +08:00 +08:00:00 +08 +0800 +080000 +08:00 +08:00:00",
		},
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

// func TestStrfTime(t *testing.T) {
// 	tt, err := timefmt.Parse("2020/07/24 09:07:29", "%Y/%m/%d %H:%M:%S")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(tt) // 2020-07-24 09:07:29 +0000 UTC

// 	str := timefmt.Format(tt, "%Y/%m/%d %H:%M:%S")
// 	fmt.Println(str) // 2020/07/24 09:07:29

// 	str = timefmt.Format(tt, "%a, %d %b %Y %T %z")
// 	fmt.Println(str) // Fri, 24 Jul 2020 09:07:29 +0000
// }

func zone(zone string) time.Time {
	location := location(zone)
	return time.Date(2021, 1, 1, 1, 1, 1, 111111111, location)
}

func location(zone string) *time.Location {
	loc, _ := time.LoadLocation(zone)
	return loc
}
