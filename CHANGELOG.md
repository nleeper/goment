# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.4.4] - 2022-01-28
- `add indonesian language support #47` from dimasdanz

## [1.4.3] - 2022-01-22
- `add brazilian portuguese language support #46` from alexandrepossebom

## [1.4.2] - 2021-06-08
- Bugfix on `IsSame() and similar functions #40` from jftuga
- IsSame and similar functions will now work if an object supplied is either a Goment struct or a pointer to a Goment struct.

## [1.4.1] - 2021-04-11
### Changed
- Bugfix on `StartOf("isoWeek") not correct on Sunday #38` from kesyn

## [1.4.0] - 2020-09-04
### Added
- Added MonthByNumber method to return the month name by number.
- Added MonthShortByNumber method to return the short month name by number.
- Added WeekdayByNumber method to return the weekday name by number. This method is locale aware.

### Changed
- Updated Weekdays method to be locale aware if bool parameter is provided.
- Updated WeekdaysShort method to be locale aware if bool parameter is provided.
- Updated WeekdaysMin method to be locale aware if bool parameter is provided.
- Support for parsing timestamps with Z timezone - https://github.com/nleeper/goment/pull/36

## [1.3.0] - 2020-08-20
### Added
- Implemented Weekday & SetWeekday methods, both are locale aware.
- Implemented Week & SetWeek methods, both are locale aware.
- Implemented WeekYear & SetWeekYear methods, both are locale aware.
- Implemented WeeksInYear method, which is locale aware.
- Implemented ISOWeeksInYear method.
- Implemented SetISOWeek method.
- Implemented SetISOWeekYear method.
- Added ability to pass day name into SetDay method, day name is locale aware. 
- Added support for weekday & weekyear tokens in parsing & formatting: e, E, w, ww, W, WW, gg, GG, gggg, GGGG
- Added support for min weekday name token in parsing: dd
- Added support for StartOf("week") && EndOf("week"), both are locale aware.

### Changed
- Updated SetDay to handle values that are outside of 0-6. Values will be treated as overflow to the current value.
- Updated the constructor that takes a DateTime object to default to the current day's values when not provided for year, month & date. For example, if the constructor is called as `goment.New(DateTime{ Year: 2010 })` and today's date is `2020-08-08`, the created Goment object will be for `2010-08-08 00:00:00`.
- Updated parsing to return an error if the weekday supplied in the parsing string does not match the weekday parsed from date fields. 
```
_, err := New("Wed 08-10-2017", "ddd MM-DD-YYYY") // 8-10-2017 is a Thursday
assert.EqualError(err, "There is a mismatch between parsed weekday and expected weekday")
```

### Removed

## [1.2.0] - 2020-07-29
### Added
- Support for locales when parsing Goment datetimes. Locales now can be passed in to the `New` function like `goment.New("sábado abr 11 22:52:51 2009", "dddd MMM DD HH:mm:ss YYYY", "es")`. Refer to the `Parsing` section in the `README.md` for more information.

## [1.1.1] - 2020-07-14
### Added
- Support for `zzzz` formatting token to display timezone name, e.g. `Central Standard Time`.

## [1.1.0] - 2020-07-13
### Added
- Support for internationalization using locales when displaying Goment datetimes. Locales are currently only supported for the `Format`, `From`, `To`, `FromNow`, `ToNow` & `Calendar` methods.
- Initial supported locales are `en`, `es` and `fr`. There is a framework setup to add new languages by adding a new file in the `/locales` folder. Refer to the `i18n` section in the `README.md` for more information.
- Added `Weekdays`, `WeekdaysShort`, `WeekdaysMin`, `Months` & `MonthsShort` methods for getting locale-specific values.
- Support for `YYYYY`, `YYYYYY` & `x` datetime formats.

### Changed
- `Format` now uses functions to replace values in layout, instead of converting to a Go datetime layout and using the Time.Format method.

### Removed
- Removed internal submodules.

## [1.0.0] - 2020-06-25
### Added
- This is the first tagged release of Goment. It has support for Go modules using `go mod`.
