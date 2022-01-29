![Go](https://github.com/nleeper/goment/workflows/Go/badge.svg)

# Goment
###### Current Version: 1.4.4
###### [Changelog](CHANGELOG.md)

Goment is a port of the popular Javascript datetime library [Moment.js](https://momentjs.com/). It follows the Moment.js API closely, with some changes to make it more Go-like (e.g. using nanoseconds instead of milliseconds). 

Goment is still a work in progress. Please feel free to fork and contribute missing methods, locale/languages functionality, or just provide more idiomatic Go if you see some areas to improve. I have a list of things that need added/fixed in [TODO.md](TODO.md), but will create issues for them at some point.

## Features
* [Parsing](#parsing) 
* [Get+Set](#get-set)
* [Manipulate](#manipulate)
* [Display](#display)
* [Query](#query)
* [i18n](#i18n)

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

Goment's parser is strict, and defaults to being accurate over forgiving.

##### Year, month, and day tokens
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
| Unix Timestamp | X | 1360013296 |
| | | |
```
goment.New("12-25-1995", "MM-DD-YYYY")
```

##### Hour, minute, second, and offset tokens
|   | Token | Output |
| - | ----- | ------ |
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
| | | |

##### Week year, week, and weekday tokens
For these, the lowercase tokens use the locale aware week start days, and the uppercase tokens use the ISO week date start days.

|   | Token | Output |
| - | ----- | ------ |
| Locale 4 digit week year | gggg | 2014 |
| Locale 2 digit week year | gg | 14 |
| Locale week of year | w ww | 1..53 |
| Locale day of week | e | 0..6 |
| Day name in locale set by Goment.SetLocale() | dd ddd dddd | Mo..Mon..Sunday |
| ISO 4 digit week year | GGGG | 2014 |
| ISO 2 digit week year | GG | 14 |
| ISO week of year | W WW | 1..53 |
| ISO day of week | E | 1..7 |
| | | |

#### From string + format + locale
As of Goment 1.2.0, a locale can now be supplied to parse locale-specific dates and times.

##### Locale-aware formats
|   | Token | Output |
| - | ----- | ------ |
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
| | | |
```
goment.New("2 de septiembre de 1999 12:30", "LLL", "es")
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
g.Get('hours') // 22
```
#### Nanosecond
Gets the nanoseconds of the Goment object.
```
g.Nanosecond() // 600
```
#### Millisecond
Gets the milliseconds of the Goment object.
```
g.Millisecond() // 330
```
#### Second
Gets the seconds of the Goment object.
```
g.Second() // 33
```
#### Minute
Gets the minutes of the Goment object.
```
g.Minute() // 45
```
#### Hour
Gets the hours of the Goment object.
```
g.Hour() // 22
```
#### Day of Month
Gets the day of the month of the Goment object.
```
g.Date() // 19
```
#### Day of Week
Gets the day of the week (Sunday = 0...) of the Goment object.
```
g.Day() // 2
```
#### Day of Week (Locale Aware)
Gets the day of the week according to the locale of the Goment object.
```
g.Weekday() // 3
```
#### ISO Day of Week
Gets the Goment object ISO day of the week with 1 being Monday and 7 being Sunday.
```
g.ISOWeekday() // 4
```
#### Day of Year
Gets the day of the year of the Goment object.
```
g.DayOfYear() // 100
```
#### Week of Year
Gets the localized week of the year for the Goment object.

The week of the year varies depending on which day is the first day of the week (Sunday, Monday, etc), and which week is the first week of the year.

For example, in the United States, Sunday is the first day of the week. The week with January 1st in it is the first week of the year.

In France, Monday is the first day of the week, and the week with January 4th is the first week of the year.

The output of `g.Week()` will depend on the locale for the Goment object.
```
g.Week() // 52
```
#### Week of Year (ISO)
Gets the ISO week of the year of the Goment object.
```
g.ISOWeek() // 6
```
#### Month
Gets the month (January = 1...) of the Goment object.
```
g.Month() // 2
```
#### Quarter
Gets the quarter (1 to 4) of the Goment object.
```
g.Quarter() // 1
```
#### Year
Gets the year of the Goment object.
```
g.Year() // 2013
```
#### Week Year
Gets the week-year according to the Goment object's locale.

Because the first day of the first week does not always fall on the first day of the year, sometimes the week-year will differ from the month year.

For example, in the US, the week that contains Jan 1 is always the first week. In the US, weeks also start on Sunday. If Jan 1 was a Monday, Dec 31 would belong to the same week as Jan 1, and thus the same week-year as Jan 1. Dec 30 would have a different week-year than Dec 31.
```
g.WeekYear() // 2012
```
#### Week Year (ISO)
Gets the ISO week-year of the Goment object.
```
g.ISOWeekYear() // 2013
```
#### Weeks In Year
Gets the number of weeks according to locale in the current Goment's year.
```
g.WeekYear() // 53
```
#### Weeks In Year (ISO)
Gets the number of weeks in the current Goment's year, according to ISO weeks.
```
g.ISOWeeksInYear() // 52
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
#### Set Nanosecond
Sets the nanoseconds for the Goment object.
```
g.SetNanosecond(60000)
```
#### Set Millisecond
Sets the milliseconds for the Goment object.
```
g.SetMillisecond(5000)
```
#### Set Second
Sets the seconds for the Goment object.
```
g.SetSecond(55)
```
#### Set Minute
Sets the minutes for the Goment object.
```
g.SetMinute(15)
```
#### Set Hour
Sets the hours for the Goment object.
```
g.SetHour(5)
```
#### Set Day of Month
Sets the day of the month for the Goment object. If the date passed in is greater than the number of days in the month, then the day is set to the last day of the month.
```
g.SetDate(21)
```
#### Set Day of Week
Sets the day of the week (Sunday = 0...) for the Goment object.
```
g.SetDay(1)
```
As of 1.3.0, a day name is also supported. This is parsed in the Goment object's locale.
```
g.SetDay("lunes")
```
#### Set Day of Week (Locale Aware)
Sets the day of the week according to the locale of the Goment object.

If the locale assigns Monday as the first day of the week, `g.SetWeekday(0)` will be Monday. If Sunday is the first day of the week, `g.SetWeekday(0)` will be Sunday.
```
g.SetWeekday(0)
```
#### Set Day of Week (ISO)
Sets the Goment object ISO day of the week with 1 being Monday and 7 being Sunday.
```
g.SetISOWeekday(2)
```
#### Set Day of Year
Sets the day of the year for the Goment object. For non-leap years, 366 is treated as 365.
```
g.SetDayOfYear(100)
```
#### Set Week of Year
Sets the localized week of the year for the Goment object. When setting the week of the year, the day of the week is retained.
```
get.SetWeek(50)
```
#### Set Week of Year (ISO)
Sets the ISO week of the year for the Goment object.

When setting the week of the year, the day of the week is retained.
```
g.SetISOWeek(52)
```
#### Set Month
Sets the month (January = 1...) of the Goment object. If new month has less days than current month, the date is pinned to the end of the target month.
```
g.SetMonth(3)
```
#### Set Quarter
Sets the quarter (1 to 4) for the Goment object.
```
g.SetQuarter(2)
```
#### Set Year
Sets the year for the Goment object.
```
g.SetYear(2010)
```
#### Set Week Year
Sets the week-year according to the Goment object's locale.
```
g.SetWeekYear(2014)
```
#### Set Week Year (ISO)
Sets the ISO week-year of the Goment object.
```
g.SetISOWeekYear(2018)
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
g.UTCOffset() // -6
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
| Quarter | Q | 1 2 3 4 |
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
| | YYYYYY | -001970 -001971 ... +001970 +001971 |
| | YYYYY | 01970 01971 ... 02010 02100 |
| | YYYY | 1970 1971 ... 2029 2030 |
| | Y | 1970 1971 ... 9999 +10000 +10001 |
| Week Year | gg | 70 71 ... 29 30 |
| | gggg | 1970 1971 ... 2029 2030 |
| Week Year (ISO) |	GG | 70 71 ... 29 30 
| | GGGG | 1970 1971 ... 2029 2030 |
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
| | zzzz | Eastern Standard Time |
| | Z | -07:00 -06:00 ... +06:00 +07:00 |
| | ZZ | -0700 -0600 ... +0600 +0700 |
| Unix Timestamp | X | 1360013296 |
| Unix Millisecond Timestamp | x | 1360013296123 |
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
g.Format('YYYY-MM-DD') // 2020-05-01
```
#### FromNow
FromNow returns the relative time from now to the Goment time.
```
g.FromNow() // 10 months ago
```
#### ToNow
ToNow returns the relative time to now to the Goment time.
```
g.ToNow() // minutes
```
#### From
From returns the relative time from the supplied time to the Goment time.
```
g.From(goment.New()) // a day ago
```
#### To
To returns the relative time from the Goment time to the supplied time.
```
g.To(goment.New()) // in a minute
```
#### Calendar
Calendar displays time relative to a given referenceTime (defaults to now).
```
g.Calendar() // Today at 1:00 PM
```
#### Difference
Diff returns the difference between two Goments as an integer.
```
g.Diff(goment.New(), 'years') // 3
```
#### ToUnix
ToUnix returns the Unix timestamp (the number of seconds since the Unix Epoch).
```
g.ToUnix() // 1360310950
```
#### DaysInMonth
DaysInMonth returns the number of days in the set month.
```
g.DaysInMonth() // 28
```
#### ToTime
ToTime returns the time.Time object that is wrapped by Goment.
```
g.ToTime()
```
#### ToArray
ToArray returns an array that mirrors the parameters from time.Date().
```
g.ToArray() // [2013 2 8 8 9 10 0]
```
#### ToDateTime
ToDateTime returns a Goment.DateTime struct.
```
g.ToDateTime() // {2013 2 8 8 9 10 0 UTC}
```
#### ToString
ToString returns an English string representation of the Goment time.
```
g.ToString() // 2006-01-02 15:04:05.999999999 -0700 MST
```
#### ToISOString
ToISOString returns a ISO8601 standard representation of the Goment time.
```
g.ToISOString() // 2016-04-12T19:46:47.286Z
```

### Query
#### IsBefore
IsBefore will check if a Goment is before another Goment. Works with Goment struct or pointer to struct.
```
g.IsBefore(goment.New()) // true
```
#### IsAfter
IsAfter will check if a Goment is after another Goment. Works with Goment struct or pointer to struct.
```
g.IsAfter(goment.New()) // false
```
#### IsSame
IsSame will check if a Goment is the same as another Goment. Works with Goment struct or pointer to struct. 
```
g.IsSame(goment.New()) // true
```
#### IsSameOrBefore
IsSameOrBefore will check if a Goment is before or the same as another Goment. Works with Goment struct or pointer to struct.
```
g.IsSameOrBefore(goment.New()) // true
```
#### IsSameOrAfter
IsSameOrAfter will check if a Goment is after or the same as another Goment. Works with Goment struct or pointer to struct.
```
g.IsSameOrAfter(goment.New()) // false
```
#### IsBetween
IsBetween will check if a Goment is between two other Goments. Works with Goment struct or pointer to struct.
```
g.IsBetween(goment.New(), goment.New().Add(5, 'days)) // true
```
#### IsDST
IsDST checks if the Goment is in daylight saving time.
```
g.IsDST() // true
```
#### IsLeapYear
IsLeapYear returns true if the Goment's year is a leap year, and false if it is not.
```
g.IsLeapYear() // false
```
#### IsTime
IsTime will check if a variable is a time.Time object.
```
g.IsTime(time.Now()) // true
```
#### IsGoment
IsGoment will check if a variable is a Goment object. Works with Goment struct or pointer to struct.
```
g.IsGoment(goment.New()) // true
```

### i18n
Goment has support for internationalization. 

In addition to assigning a global locale, you can assign a locale to a specific Goment object.

\* Currently, only formatting functions like `Format`, `To`, `From`, `ToNow`, `FromNow` & `Calendar` use locales. Only English (United States) datetime formats are able to be parsed at this time.

#### Changing global locale
By default, Goment uses English (United States) locale strings. Changing the global locale does not affect existing Goment instances.
```
SetLocale("es")
```

#### Getting global locale
```
Locale() // es
```

#### Changing locale for Goment instance
```
g.SetLocale("fr")
```

#### Getting locale for Goment instance
```
g.Locale() // fr
```

#### Weekdays
Returns a list of weekdays in the current locale. If the parameter provided is a bool value of true,
the list returned will take the current locale's first day of week into consideration. If no true parameter
is provided, the first day will always be Sunday for the locale.
```
g.SetLocale("en")
g.Weekdays() // [ "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday" ]

g.SetLocale("es")
g.Weekdays(true) // ["lunes", "martes", "miércoles", "jueves", "viernes", "sábado", "domingo"]
g.Weekdays() // ["domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"]
```
#### WeekdaysShort
Returns a list of abbreviated weekdays in the current locale. If the parameter provided is a bool value of true, the list returned will take the current locale's first day of week into consideration. If no true parameter is provided, the first day will always be Sunday for the locale.
```
g.SetLocale("en")
g.WeekdaysShort() // [ "Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat" ]

g.SetLocale("es")
g.WeekdaysShort(true) // ["lun.", "mar.", "mié.", "jue.", "vie.", "sáb.", "dom."]
g.WeedaysShort() // ["dom.", "lun.", "mar.", "mié.", "jue.", "vie.", "sáb."]
```
#### WeekdaysMin
Returns a list of abbreviated weekdays in the current locale. If the parameter provided is a bool value of true, the list returned will take the current locale's first day of week into consideration. If no true parameter is provided, the first day will always be Sunday for the locale.
```
g.SetLocale("en")
g.WeekdaysMin() // [ "Su", "Mo", "Tu", "We", "Th", "Fr", "Sa" ]

g.SetLocale("es")
g.WeekdaysMin(true) // ["lu", "ma", "mi", "ju", "vi", "sá", "do"]
g.WeedaysMin() // ["do", "lu", "ma", "mi", "ju", "vi", "sá"]
```
#### WeekdayByNumber
Returns the weekday name by number in the current locale. If the first parameter is a bool value of true,
the name returned will take the current locale's first day of week into consideration. If no true parameter is provided, the first day will always be Sunday for the locale.
```
g.SetLocale("en")
g.WeekdayByNumber(0) // "Sunday"
g.WeekdayByNumber(true, 0) // "Sunday"

g.SetLocale("es")
g.WeekdayByNumber(0) // "domingo"
g.WeekdayByNumber(true, 0) // "lunes"
```
#### Months
Returns a list of months in the current locale.
```
g.Months() // [ "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December" ]
```
#### MonthByNumber
Returns the month name by number in the current locale.
```
g.MonthByNumber(1) // "January"
```
#### MonthsShort
Returns a list of abbreviated month names in the current locale.
```
g.MonthsShort() // [ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec" ]
```
#### MonthShortByNumber
Returns the abbreviated month name by number in the current locale.
```
g.MonthShortByNumber(2) // "Feb"
```

#### Adding a new locale
To add a new locale, there are a few steps to follow. You must first add a new file in the `/locales` folder. This should be named the locale code, e.g. `fr.go`. Inside this file, you need to create a new `LocaleDetails` object and provide the required values for month names, weekday names, ordinal function, etc. Please use one of the existing locales for reference.

After you've created the locale file, add a line to `locale.go` in the `supportedLocales` map. This should be a map from the locale code to an instance of the `LocaleDetails` object you created above.

Lastly, please add test cases to `locale_test.go` that test the different datetime formats, and the relative time formats.