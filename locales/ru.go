package locales

import (
	"strconv"
	"strings"
)

// RuLocale is the Russian language locale.
var RuLocale = newLocale(
	"ru",
	strings.Split("Воскресенье_Понедельник_Вторник_Среда_Четверг_Пятница_Суббота", "_"),
	strings.Split("Вс_Пн_Вт_Ср_Чт_Пт_Сб", "_"),
	strings.Split("Вс_Пн_Вт_Ср_Чт_Пт_Сб", "_"),
	strings.Split("Январь_Февраль_Март_Апрель_Май_Июнь_Июль_Август_Сентябрь_Октябрь_Ноябрь_Декабрь", "_"),
	strings.Split("Янв_Фев_Мар_Апр_Май_Июн_Июл_Авг_Сен_Окт_Ноя_Дек", "_"),
	func(num int, period string) string {
		suffix := "th"
		switch period {
		case "M":
		case "d":
		case "DDD":
			suffix = "-й"
		case "D":
			suffix = "-го"
		case "w":
		case "W":
			suffix = "-я"
		}
		return strconv.Itoa(num) + suffix
	},
	nil,
	week{Dow: 1, Doy: 4},
	longDateFormats{
		"LTS":  "H:mm:ss",
		"LT":   "H:mm",
		"L":    "DD.MM.YYYY",
		"LL":   "D MMMM YYYY г.",
		"LLL":  "D MMMM YYYY г., H:mm",
		"LLLL": "dddd, D MMMM YYYY г., H:mm",
	},
	relativeTimeFormats{
		"future": "через %s",
		"past":   "%s назад",
		"s":      "несколько секунд",
		"ss":     "%d секунд",
		"m":      "минуту",
		"mm":     "%d минут",
		"h":      "час",
		"hh":     "%d часов",
		"d":      "день",
		"dd":     "%d дней",
		"M":      "месяц",
		"MM":     "%d месяцев",
		"y":      "год",
		"yy":     "%d года",
	},
	calendarFunctions{
		"sameDay": func(hours int, day int) string {
			return "[Сегодня в] LT"
		},
		"nextDay": func(hours int, day int) string {
			return "[Завтра в] LT"
		},
		"nextWeek": func(hours int, day int) string {
			return "dddd [в сл.] LT"
		},
		"lastDay": func(hours int, day int) string {
			return "[Вчера в] LT"
		},
		"lastWeek": func(hours int, day int) string {
			return "[В прошлый] dddd [в] LT"
		},
		"sameElse": func(hours int, day int) string {
			return "L"
		},
	},
	`(?i)(Январь|Февраль|Март|Апрель|Май|Июнь|Июль|Август|Сентябрь|Октябрь|Ноябрь|Декабрь)`,
	`(?i)(Янв|Фев|Мар|Апр|Май|Июн|Июл|Авг|Сен|Окт|Ноя|Дек)`,
	`(?i)(Воскресенье|Понедельник|Вторник|Среда|Четверг|Пятница|Суббота)`,
	`(?i)(Вос|Пон|Вто|Сре|Чет|Пят|Суб)`,
	`(?i)(Вс|Пн|Вт|Ср|Чт|Пт|Сб)`,
	`\d{1,2}(й|го|я)`,
)
