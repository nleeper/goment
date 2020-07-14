# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
