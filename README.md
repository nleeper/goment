# Goment
Goment is a port of the popular Javascript datetime library [Moment.js](https://momentjs.com/). It follows the Moment.js API closely, with some changes to make it more Go-like (e.g. using nanoseconds instead of milliseconds). 

Goment is still a work in progress. Please feel free to fork and contribute missing methods, locale/languages functionality, or just provide more idiomatic Go if you see some areas to improve. I have a list of things that need added/fixed in [TODO.md](TODO.md), but will create issues for them at some point.

## Features
* [Parsing](#parsing) 
* [Get+Set](#get-set)
* [Manipulate](#manipulate)
* [Display](#display)
* [Query](#query)

### Parsing
#### From now
Creates a Goment object for the current local time returned by time.Now().
```
goment.New()
```
#### From ISO 8601 string
Creates a Goment object by parsing the string as an ISO 8601 date time. The timezone will be UTC unless supplied in the string.
```
goment.New('2013-02-08 09:30:26')
```
#### From string + format
Creates a Goment object by parsing the string using the supplied format. The timezone will be the local timezone unless supplied in the string.

The parsing tokens are similar to the formatting tokens used in [Goment#Format](#format).

##### Supported tokens
|   | Token | Output |
| - | ----- | ------ |
| Month | M | 1 2 ... 11 12 |
| | MM | 01 01 ... 11 12 |
| | MMM | Jan Feb ... Nov Dec |
| | MMMM | January February ... November December |
| Day of Month | D | 1 2 ... 30 31 |
| | Do | 1st 2nd ... 30th 31st |
| | DD | 01 02 ... 30 31 |
| Day of Year | DDD	 | 1 2 ... 364 365 |
| | DDDD | 001 002 ... 364 365 |
| Year | YY | 70 71 ... 29 30 |
| | YYYY | 1970 1971 ... 2029 2030 |
| | Y | -25 |
| Quarter | Q | 1 2 3 4 |
| AM/PM	| A | AM PM |
| | a |	am pm |
| Hour| H | 0 1 ... 22 23 |
| | HH | 00 01 ... 22 23 |
| | h | 1 2 ... 11 12 |
| | hh | 01 02 ... 11 12 |
| | k | 1 2 ... 23 24 |
| | kk | 01 02 ... 23 24 |
| Minute | m | 0 1 ... 58 59 |
| | mm | 00 01 ... 58 59 |
| Second | s | 0 1 ... 58 59 |
| | ss | 00 01 ... 58 59 |
| Time Zone	| Z | -07:00 -06:00 ... +06:00 +07:00 |
| | ZZ | -0700 -0600 ... +0600 +0700 |
| Unix Timestamp | X | 1360013296 |
```
goment.New("12-25-1995", "MM-DD-YYYY")
```

#### From Unix nanoseconds
Creates a Goment object from the Unix nanoseconds since the Unix Epoch.
```
goment.New(time.Now().UnixNano())
```
#### From Unix seconds
Creates a Goment object from the Unix timestamp (seconds since the Unix Epoch).
```
goment.Unix(1318781876)
```
#### From [Go Time object](https://golang.org/pkg/time/#Time)
Creates a Goment object from the supplied Go time object.
```
goment.New(time.Date(2015, 11, 10, 5, 30, 0, 0, time.UTC))
```
#### From a Goment clone
Creates a Goment object from a clone of the supplied Goment object.
```
goment.New(goment.New('2011-05-08'))
```
#### From Goment DateTime object
Creates a Goment object from a Goment DateTime object.
``` 
goment.New(DateTime{
    Year:  2015,
    Month: 1,
    Day:   25,
    Hour:  10,
})
```

### Get+Set
#### Get
Get is a string getter using the supplied units.

##### Supported units
* y, year, years
* M, month, months
* D, date, dates
* h, hour, hours
* m, minute, minutes
* s, second, seconds
* ms, millisecond, milliseconds
* ns, nanosecond, nanoseconds

```
g.Get('hours')
```
#### Nanosecond
Get the nanoseconds of the Goment object.
```
g.Nanosecond()
```
#### Millisecond
Get the milliseconds of the Goment object.
```
g.Millisecond()
```
#### Second
Get the seconds of the Goment object.
```
g.Second()
```
#### Minute
Get the minutes of the Goment object.
```
g.Minute()
```
#### Hour
Get the hours of the Goment object.
```
g.Hour()
```
#### Date
Get the day of the month of the Goment object.
```
g.Date()
```
#### Day
Get the day of the week (Sunday = 0...) of the Goment object.
```
g.Day()
```
#### ISOWeekday
Gets the Goment object ISO day of the week with 1 being Monday and 7 being Sunday.
```
g.ISOWeekday()
```
#### DayOfYear
Gets the day of the year of the Goment object.
```
g.DayOfYear()
```
#### ISOWeek
Gets the ISO week of the year of the Goment object.
```
g.ISOWeek()
```
#### Month
Gets the month (January = 1...) of the Goment object.
```
g.Month()
```
#### Quarter
Gets the quarter (1 to 4) of the Goment object.
```
g.Quarter()
```
#### Year
Gets the year of the Goment object.
```
g.Year()
```
#### ISOWeekYear
Gets the ISO week-year of the Goment object.
```
g.ISOWeekYear()
```
#### Set
Set is a generic setter, accepting units as the first argument, and value as the second.

##### Supported units
* y, year, years
* M, month, months
* D, date, dates
* h, hour, hours
* m, minute, minutes
* s, second, seconds
* ms, millisecond, milliseconds
* ns, nanosecond, nanoseconds

```
g.Set(6, 'hour')
```
#### SetNanosecond
Set the nanoseconds for the Goment object.
```
g.SetNanosecond(60000)
```
#### SetMillisecond
Set the milliseconds for the Goment object.
```
g.SetMillisecond(5000)
```
#### SetSecond
Set the seconds for the Goment object.
```
g.SetSecond(55)
```
#### SetMinute
Set the minutes for the Goment object.
```
g.SetMinute(15)
```
#### SetHour
Set the hours for the Goment object.
```
g.SetHour(5)
```
#### SetDate
Set the day of the month for the Goment object. If the date passed in is greater than the number of days in the month, then the day is set to the last day of the month.
```
g.SetDate(21)
```
#### SetDay
Set the day of the week (Sunday = 0...) for the Goment object.
```
g.SetDay(1)
```
#### SetISOWeekday
Sets the Goment object ISO day of the week with 1 being Monday and 7 being Sunday.
```
g.SetISOWeekday(2)
```
#### SetDayOfYear
Sets the day of the year for the Goment object. For non-leap years, 366 is treated as 365.
```
g.SetDayOfYear(100)
```
#### SetMonth
Sets the month (January = 1...) of the Goment object. If new month has less days than current month, the date is pinned to the end of the target month.
```
g.SetMonth(3)
```
#### SetQuarter
Sets the quarter (1 to 4) for the Goment object.
```
g.SetQuarter(2)
```
#### SetYear
Sets the year for the Goment object.
```
g.SetYear(2010)
```

### Manipulate
#### Add
Add mutates the Goment object by adding time. The first argument can either be a time.Duration, or an integer representing the number of the unit to add. The second argument should be a unit.

##### Supported units
* y, year, years
* Q, quarter, quarters
* M, month, months
* w, week, weeks
* d, day, days
* h, hour, hours
* m, minute, minutes
* s, second, seconds
* ms, millisecond, milliseconds
* ns, nanosecond, nanoseconds

```
g.Add(1, 'days')
```
#### Subtract
Subtract mutates the Goment object by subtracting time. The first argument can either be a time.Duration, or an integer representing the number of the unit to add. The second argument should be a unit.

##### Supported units
* y, year, years
* Q, quarter, quarters
* M, month, months
* w, week, weeks
* d, day, days
* h, hour, hours
* m, minute, minutes
* s, second, seconds
* ms, millisecond, milliseconds
* ns, nanosecond, nanoseconds

```
g.Subtract(5, 'hours')
```
#### StartOf
StartOf mutates the Goment object by setting it to the start of a unit of time.

##### Supported units
* y, year, years
* Q, quarter, quarters
* M, month, months
* w, week, weeks
* W, isoWeek, isoWeeks
* d, day, days
* h, hour, hours
* m, minute, minutes
* s, second, seconds

```
g.StartOf('day')
```
#### EndOf
EndOf mutates the Goment object by setting it to the end of a unit of time.
##### Supported units
* y, year, years
* Q, quarter, quarters
* M, month, months
* w, week, weeks
* W, isoWeek, isoWeeks
* d, day, days
* h, hour, hours
* m, minute, minutes
* s, second, seconds

```
g.EndOf('month')
```
#### Local
Local will set the Goment to use local time.
```
g.Local()
```
#### UTC
UTC will set the Goment to use UTC time.
```
g.UTC()
```
#### UTCOffset
UTCOffset gets the Goment's UTC offset in minutes.
```
g.UTCOffset()
```
#### SetUTCOffset
SetUTCOffset sets the Goment's UTC offset in minutes. If the offset is less than 16 and greater than -16, the value is treated as hours.
```
g.SetUTCOffset(120)
```

### Display
#### Format
Format takes a string of tokens and replaces them with their corresponding values to display the Goment.

##### Supported tokens
|   | Token | Output |
| - | ----- | ------ |
| Month | M | 1 2 ... 11 12 |
| | Mo | 1st 2nd ... 11th 12th |
| | MM | 01 01 ... 11 12 |
| | MMM | Jan Feb ... Nov Dec |
| | MMMM | January February ... November December |
| Day of Month | D | 1 2 ... 30 31 |
| | Do | 1st 2nd ... 30th 31st |
| | DD | 01 02 ... 30 31 |
| Day of Year | DDD	 | 1 2 ... 364 365 |
| | DDDo | 1st 2nd ... 364th 365th |
| | DDDD | 001 002 ... 364 365 |
| Day of Week | d | 0 1 ... 5 6 |
| | do | 0th 1st ... 5th 6th |
| | dd | Su Mo ... Fr Sa |
| | ddd | Sun Mon ... Fri Sat |
| | dddd | Sunday Monday ... Friday Saturday |
| Day of Week (Locale) | e | 0 1 ... 5 6 |
| Day of Week (ISO) | E | 1 2 ... 6 7 |
| Week of Year | w | 1 2 ... 52 53 |
| | wo | 1st 2nd ... 52nd 53rd |
| | ww | 01 02 ... 52 53 |
| Week of Year (ISO) | W | 1 2 ... 52 53 |
| | Wo | 1st 2nd ... 52nd 53rd |
| | WW | 01 02 ... 52 53 |
| Year | YY | 70 71 ... 29 30 |
| | YYYY | 1970 1971 ... 2029 2030 |
| | Y | 1970 1971 ... 9999 +10000 +10001 |
| Quarter | Q | 1 2 3 4 |
| AM/PM	| A | AM PM |
| | a |	am pm |
| Hour| H | 0 1 ... 22 23 |
| | HH | 00 01 ... 22 23 |
| | h | 1 2 ... 11 12 |
| | hh | 01 02 ... 11 12 |
| | k | 1 2 ... 23 24 |
| | kk | 01 02 ... 23 24 |
| Minute | m | 0 1 ... 58 59 |
| | mm | 00 01 ... 58 59 |
| Second | s | 0 1 ... 58 59 |
| | ss | 00 01 ... 58 59 |
| Time Zone	| z or zz | EST CST ... MST PST |
| | Z | -07:00 -06:00 ... +06:00 +07:00 |
| | ZZ | -0700 -0600 ... +0600 +0700 |
| Unix Timestamp | X | 1360013296 |
| Time | LT | 8:30 PM |
| Time with seconds	| LTS | 8:30:25 PM |
| Month numeral, day of month, year	| L	| 09/04/1986 |
| | l | 9/4/1986 |
| Month name, day of month, year | LL | September 4, 1986 |
| | ll | Sep 4, 1986 |
| Month name, day of month, year, time | LLL | September 4, 1986 8:30 PM |
| | lll	| Sep 4, 1986 8:30 PM |
| Month name, day of month, day of week, year, time	| LLLL |	Thursday, September 4, 1986 8:30 PM |
| | llll | Thu, Sep 4, 1986 8:30 PM |

```
g.Format('YYYY-MM-DD')
```
#### FromNow
FromNow returns the relative time from now to the Goment time.
```
g.FromNow()
```
#### ToNow
ToNow returns the relative time to now to the Goment time.
```
g.ToNow()
```
#### From
From returns the relative time from the supplied time to the Goment time.
```
g.From(goment.New())
```
#### To
To returns the relative time from the Goment time to the supplied time.
```
g.To(goment.New())
```
#### Calendar
Calendar displays time relative to a given referenceTime (defaults to now).
```
g.Calendar()
```
Difference
Diff returns the difference between two Goments as an integer.
```
g.Diff(goment.New(), 'years')
```
#### ToUnix
ToUnix returns the Unix timestamp (the number of seconds since the Unix Epoch).
```
g.ToUnix()
```
#### DaysInMonth
DaysInMonth returns the number of days in the set month.
```
g.DaysInMonth()
```
#### ToTime
ToTime returns the time.Time object that is wrapped by Goment.
```
g.ToTime()
```
#### ToArray
ToArray returns an array that mirrors the parameters from time.Date().
```
g.ToArray()
```
#### ToDateTime
ToDateTime returns a Goment.DateTime struct.
```
g.ToDateTime()
```
#### ToString
ToString returns a string representation of the Goment time.
```
g.ToString()
```
#### ToISOString
ToISOString returns a ISO8601 standard representation of the Goment time.
```
g.ToISOString()
```

### Query
#### IsBefore
IsBefore will check if a Goment is before another Goment.
```
g.IsBefore(goment.New())
```
#### IsAfter
IsAfter will check if a Goment is after another Goment.
```
g.IsAfter(goment.New())
```
#### IsSame
IsSame will check if a Goment is the same as another Goment.
```
g.IsSame(goment.New())
```
#### IsSameOrBefore
IsSameOrBefore will check if a Goment is before or the same as another Goment.
```
g.IsSameOrBefore(goment.New())
```
#### IsSameOrAfter
IsSameOrAfter will check if a Goment is after or the same as another Goment.
```
g.IsSameOrAfter(goment.New())
```
#### IsBetween
IsBetween will check if a Goment is between two other Goments.
```
g.IsBetween(goment.New(), goment.New().Add(5, 'days))
```
#### IsDST
IsDST checks if the Goment is in daylight saving time.
```
g.IsDST()
```
#### IsLeapYear
IsLeapYear returns true if the Goment's year is a leap year, and false if it is not.
```
g.IsLeapYear()
```
#### IsTime
IsTime will check if a variable is a time.Time object.
```
g.IsTime(time.Now())
```
#### IsGoment
IsGoment will check if a variable is a Goment object.
```
g.IsGoment(goment.New())
```






