# Time format

## Option 1: Go standard library

```
t.Format("20060102150405")
```

| Unit                                 | Golang Layout | Examples                                           | Note                          |
| ------------------------------------ | ------------- | -------------------------------------------------- | ----------------------------- |
| Month                                | January       | January, February, December                        |                               |
| Month                                | Jan           | Jan, Feb, Dec                                      |                               |
| Month                                | 1             | 1, 2, 12                                           |                               |
| Month                                | 01            | 01, 02, 12                                         |                               |
| Day                                  | Monday        | Monday, Wednesday, Sunday                          |                               |
| Day                                  | Mon           | Mon, Wed, Sun                                      |                               |
| Day                                  | 2             | 1, 2, 11, 31                                       |                               |
| Day                                  | 02            | 01, 02, 11, 31                                     | zero padded day of the month  |
| Day                                  | 002           | 001, 002, 011, 031, 145, 365, 366                  | zero padded day of the year   |
| Day                                  | \_2           | " 1", " 2", "11", "31"                             | space padded day of the month |
| Day                                  | \_\_2         | "&nbsp; 1", "&nbsp; 2", " 11", " 31", "365", "366" | space padded day of the year  |
| Hour 24h                             | 15            | 00, 01, 12, 23                                     |                               |
| Hour 12h                             | 3             | 1, 2, 12                                           |                               |
| Hour 12h                             | 03            | 01, 02, 12                                         |                               |
| Minute                               | 4             | 0, 4 ,10, 35                                       |                               |
| Minute                               | 04            | 00, 04 ,10, 35                                     |                               |
| Second                               | 5             | 0, 5, 25                                           |                               |
| Second                               | 05            | 00, 05, 25                                         |                               |
| Year                                 | 06            | 21, 81, 01                                         |                               |
| Year                                 | 2006          | 2021, 1981, 0001                                   |                               |
| Part of day                          | PM            | AM, PM                                             |                               |
| Part of day                          | pm            | am, pm                                             |                               |
| 10<sup>-1</sup> to 10<sup>-9</sup> s | .0 .000000000 | .1, .199000000                                     | Trailing zeros included       |
| 10<sup>-1</sup> to 10<sup>-9</sup> s | .9 .999999999 | .1, .199                                           | Trailing zeros omitted        |
| Time zone                            | MST           | UTC, MST, CET                                      |                               |
| Time zone                            | Z07           | Z, +08, -05                                        | Z is for UTC                  |
| Time zone                            | Z0700         | Z, +0800, -0500                                    | Z is for UTC                  |
| Time zone                            | Z070000       | Z, +080000, -050000                                | Z is for UTC                  |
| Time zone                            | Z07:00        | Z, +08:00, -05:00                                  | Z is for UTC                  |
| Time zone                            | Z07:00:00     | Z, +08:00:00, -05:00:00                            | Z is for UTC                  |
| Time zone                            | -07           | +00, +08, -05                                      |                               |
| Time zone                            | -0700         | +0000, +0800, -0500                                |                               |
| Time zone                            | -070000       | +000000, +080000, -050000                          |                               |
| Time zone                            | -07:00        | +00:00, +08:00, -05:00                             |                               |
| Time zone                            | -07:00:00     | +00:00:00, +08:00:00, -05:00:00                    |                               |
| Week                                 |               | 1, 5, 15, 53                                       | year, week := ts.ISOWeek()    |

In Golang 1.17+ for fraction of seconds (.999 or .000) you can use `,` instead of `.` (,999 or ,000) but output is always with `.`!!!

## Option 2: strftime implementation

```
import strftime "github.com/itchyny/timefmt-go"
```

```
strftime.Format(t, "%Y%m%d%H%M%S")
```

https://github.com/itchyny/timefmt-go